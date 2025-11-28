package main

import (
	"fmt"
	"github.com/arvinbostani/Snyper.git/sniff"
	"github.com/arvinbostani/Snyper.git/ui"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var iface string
	var newWin bool

	rootCmd := &cobra.Command{
		Use:   "sniff",
		Short: "Snypher - Local network monitor",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && iface == "" {
				iface = args[0]
			}

			if iface == "" {
				fmt.Println("Usage: sniff -n <interface> OR sniff <interface>")
				os.Exit(1)
			}

			if newWin {
				if err := ui.OpenInNewTerminal(iface); err != nil {
					log.Fatalf("failed to open new terminal: %v", err)
				}
				return
			}

			packetChan := make(chan sniff.PacketInfo, 1024)
			go sniff.StartCapture(iface, packetChan)

			if err := ui.StartUI(packetChan); err != nil {
				log.Fatalf("ui error: %v", err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&iface, "NetMon", "n", "", "network interface to monitor")
	rootCmd.Flags().BoolVarP(&newWin, "new", "w", false, "open in new terminal window")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
