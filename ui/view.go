package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	timerLeft time.Duration
	status    string
}

func (m model) Init() tea.Cmd {
	return tick()
}

// 1秒ごとに更新
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

// 入力処理
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.timerLeft > 0 {
		m.timerLeft -= time.Second
		return m, tick()
	}
	m.status = " ☕️ Break Time!"
	return m, tea.Quit
}

func (m model) View() string {
	timerText := fmt.Sprintf("⏳ Pomodoro: %02d:%02d", int(m.timerLeft.Minutes()), int(m.timerLeft.Seconds())%60)
	// statusText := StatusStyle.Render(m.status)
	return TimerStyle.Render(timerText)
}

// TUI実行
func RunTimer(minutes int) {
	p := tea.NewProgram(model{timerLeft: time.Duration(minutes) * time.Minute})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
	}
}
