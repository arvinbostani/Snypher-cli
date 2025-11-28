package sniff

import (
	"fmt"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func ParsePacket(packet gopacket.Packet) *PacketInfo {
	network := packet.NetworkLayer()
	transport := packet.TransportLayer()

	if network == nil || transport == nil {
		return nil
	}

	src := network.NetworkFlow().Src().String()
	dst := network.NetworkFlow().Dst().String()
	proto := ""
	info := ""
	susp := false

	switch t := transport.(type) {
	case *layers.TCP:
		proto = "TCP"
		info = fmt.Sprintf("%s->%s", t.SrcPort, t.DstPort)
		if t.DstPort == 22 || t.DstPort == 3389 || t.DstPort == 5900 {
			susp = true
			info += " | suspicious-port"
		}

	case *layers.UDP:
		proto = "UDP"
		info = fmt.Sprintf("%s->%s", t.SrcPort, t.DstPort)
		if dnsLayer := packet.Layer(layers.LayerTypeDNS); dnsLayer != nil {
			dns := dnsLayer.(*layers.DNS)
			if len(dns.Questions) > 0 {
				q := string(dns.Questions[0].Name)
				info += " | DNS:" + q
			}
		}
	}

	if app := packet.ApplicationLayer(); app != nil {
		payload := string(app.Payload())
		if strings.HasPrefix(payload, "GET") || strings.HasPrefix(payload, "POST") ||
			strings.HasPrefix(payload, "PUT") || strings.HasPrefix(payload, "DELETE") {
			info += " | HTTP"
			if len(payload) > 1500 {
				susp = true
				info += " | large-payload"
			}
		}
	}

	return &PacketInfo{
		Source:      src,
		Destination: dst,
		Protocol:    proto,
		Info:        info,
		Suspicious:  susp,
	}
}
