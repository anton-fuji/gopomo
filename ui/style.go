package ui

import "github.com/charmbracelet/lipgloss"

// タイマーのスタイル
var (
	TimerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1E90FF")).
			Background(lipgloss.Color("#000000")).
			Bold(true).
			Padding(1, 2).
			Margin(1, 0)

	StatusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4682B4")).
			Italic(true)
)
