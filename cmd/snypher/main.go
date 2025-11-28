package main

import (
	"fmt"
	"github.com/arvinbostani/Snyper.git/sniff"
	"github.com/arvinbostani/Snyper.git/ui"
	"github.com/google/gopacket/pcap"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var iface string
	var newWin bool

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all available network interfaces",
		Run: func(cmd *cobra.Command, args []string) {
			listInterfaces()
		},
	}

	rootCmd := &cobra.Command{
		Use:   "snypher",
		Short: "Snypher - Local network monitor",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && iface == "" {
				iface = args[0]
			}

			if iface == "" {
				fmt.Println("❌ No network interface provided.")
				fmt.Println("Usage:")
				fmt.Println("  snypher -r <interface>")
				fmt.Println("  snypher <interface>")
				fmt.Println("  snypher list   # show interfaces")
				listInterfaces()
				os.Exit(1)
			}

			resolvedIface, ok := resolveInterface(iface)
			if !ok {
				fmt.Printf("❌ Interface '%s' not found.\n", iface)
				listInterfaces()
				os.Exit(1)
			}
			iface = resolvedIface

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

	rootCmd.Flags().StringVarP(&iface, "run", "r", "", "network interface to monitor")
	rootCmd.Flags().BoolVarP(&newWin, "new", "w", false, "open in new terminal window")

	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listInterfaces() {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println("Error listing interfaces:", err)
		return
	}

	fmt.Println("\nAvailable Network Interfaces:")

	for _, iface := range ifaces {
		name := iface.Name
		desc := iface.Description

		if runtime.GOOS == "windows" && desc != "" {
			fmt.Printf(" - %s (%s)\n", name, desc)
		} else {
			fmt.Printf(" - %s\n", name)
		}
	}
	fmt.Println()
}

func interfaceExists(name string) bool {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		return false
	}
	for _, i := range ifaces {
		if i.Name == name {
			return true
		}
	}
	return false
}

func resolveInterface(input string) (string, bool) {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		return "", false
	}

	for _, i := range ifaces {
		if i.Description == input {
			return i.Name, true
		}
	}

	inputLower := strings.ToLower(input)
	for _, i := range ifaces {
		if strings.Contains(strings.ToLower(i.Description), inputLower) {
			return i.Name, true
		}
	}

	for _, i := range ifaces {
		if i.Name == input {
			return i.Name, true
		}
	}

	return "", false
}
