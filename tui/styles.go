package tui

import "github.com/charmbracelet/lipgloss"

var (
	DefaultStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("5"))

	WarningStyle = lipgloss.NewStyle().Inherit(DefaultStyle).Foreground(lipgloss.Color("3"))
)
