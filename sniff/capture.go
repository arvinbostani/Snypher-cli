package sniff

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func StartCapture(device string, out chan<- PacketInfo) {
	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("pcap open failed: %v", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		if info := ParsePacket(packet); info != nil {
			select {
			case out <- *info:
			default:

			}
		}
	}
}
