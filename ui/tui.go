package ui

import (
	"fmt"
	"github.com/arvinbostani/Snyper.git/sniff"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	packets []sniff.PacketInfo
	ch      <-chan sniff.PacketInfo
	quit    chan struct{}
}

func NewModel(ch <-chan sniff.PacketInfo) model {
	return model{packets: []sniff.PacketInfo{}, ch: ch, quit: make(chan struct{})}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tea.KeyMsg:
		if v.String() == "ctrl+c" || v.String() == "q" {
			close(m.quit)
			return m, tea.Quit
		}
	case sniff.PacketInfo:
		m.packets = append(m.packets, v)
		if len(m.packets) > 40 {
			m.packets = m.packets[1:]
		}
	}
	return m, nil
}

func (m model) View() string {
	out := TitleStyle.Render("🛰️ Snypher") + "\n\n"
	for i := len(m.packets) - 1; i >= 0; i-- {
		p := m.packets[i]
		prefix := OkStyle.Render("[OK]")
		if p.Suspicious {
			prefix = SusStyle.Render("[SUS]")
		}
		out += fmt.Sprintf("%s %s -> %s | %s %s\n", prefix, p.Source, p.Destination, p.Protocol, MetaStyle.Render(p.Info))
	}
	out += "\nPress q or Ctrl+C to quit."
	return out
}

func StartUI(ch <-chan sniff.PacketInfo) error {
	m := NewModel(ch)
	p := tea.NewProgram(m, tea.WithAltScreen())

	go func() {
		for pkt := range ch {
			p.Send(pkt)
		}
	}()

	go func() {
		for range time.Tick(500 * time.Millisecond) {
			p.Send(tea.Msg(time.Now()))
		}
	}()

	return p.Start()
}
