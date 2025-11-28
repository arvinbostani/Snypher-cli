package sniff

type PacketInfo struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Protocol    string `json:"protocol"`
	Info        string `json:"info"`
	Suspicious  bool   `json:"suspicious"`
}
