package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFF7")).Bold(true).Underline(true)

	OkStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF9F"))
	SusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0044")).Bold(true)

	MetaStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9E9E9E"))
)
