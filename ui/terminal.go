package ui

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenInNewTerminal(iface string) error {
	cmdStr := fmt.Sprintf("sniff -n %s", iface)

	switch runtime.GOOS {
	case "linux":
		if path, _ := exec.LookPath("gnome-terminal"); path != "" {
			return exec.Command("gnome-terminal", "--", "sh", "-c", cmdStr).Start()
		}
		if path, _ := exec.LookPath("konsole"); path != "" {
			return exec.Command("konsole", "-e", "sh", "-c", cmdStr).Start()
		}
		if path, _ := exec.LookPath("xfce4-terminal"); path != "" {
			return exec.Command("xfce4-terminal", "-e", "sh", "-c", cmdStr).Start()
		}
		if path, _ := exec.LookPath("x-terminal-emulator"); path != "" {
			return exec.Command("x-terminal-emulator", "-e", "sh", "-c", cmdStr).Start()
		}
		return fmt.Errorf("no compatible terminal found")

	case "windows":
		if path, _ := exec.LookPath("wt"); path != "" {
			return exec.Command("wt", "cmd", "/k", cmdStr).Start()
		}
		return exec.Command("cmd", "/C", "start", "cmd", "/k", cmdStr).Start()

	default:
		return fmt.Errorf("unsupported OS for opening new terminal: %s", runtime.GOOS)
	}
}
