package gui

import tea "github.com/charmbracelet/bubbletea"

type Bubble = tea.Model

type NavMsg struct {
	Msg string
}

func (n NavMsg) String() string {
	return n.Msg
}

const (
	navUp    = "up"
	navDown  = "down"
	navLeft  = "left"
	navRight = "right"
	navEnter = "enter"
	navBack  = "back"
	navQuit  = "quit"
)

var DefaultKeyBinds = map[string]tea.Cmd{
	"w": func() tea.Msg {
		return NavMsg{navUp}
	},
	"up": func() tea.Msg {
		return NavMsg{navUp}
	},
	"a": func() tea.Msg {
		return NavMsg{navLeft}
	},
	"left": func() tea.Msg {
		return NavMsg{navLeft}
	},
	"s": func() tea.Msg {
		return NavMsg{navDown}
	},
	"down": func() tea.Msg {
		return NavMsg{navDown}
	},
	"d": func() tea.Msg {
		return NavMsg{navRight}
	},
	"right": func() tea.Msg {
		return NavMsg{navRight}
	},
	"enter": func() tea.Msg {
		return NavMsg{navEnter}
	},
	"esc": func() tea.Msg {
		return NavMsg{navBack}
	},
	"q": func() tea.Msg {
		return NavMsg{navQuit}
	},
	"ctrl+c": func() tea.Msg {
		return NavMsg{navQuit}
	},
}
