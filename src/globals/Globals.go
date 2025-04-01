package globals

import tea "github.com/charmbracelet/bubbletea"

type NavMsg struct {
	Msg string
}

func (n NavMsg) String() string {
	return n.Msg
}

//region navigationDefaults

const (
	NavUp    = "up"
	NavDown  = "down"
	NavLeft  = "left"
	NavRight = "right"
	NavEnter = "enter"
	NavBack  = "back"
	NavQuit  = "quit"
)

var DefaultKeyBinds = map[string]tea.Cmd{
	"w": func() tea.Msg {
		return NavMsg{NavUp}
	},
	"up": func() tea.Msg {
		return NavMsg{NavUp}
	},
	"a": func() tea.Msg {
		return NavMsg{NavLeft}
	},
	"left": func() tea.Msg {
		return NavMsg{NavLeft}
	},
	"s": func() tea.Msg {
		return NavMsg{NavDown}
	},
	"down": func() tea.Msg {
		return NavMsg{NavDown}
	},
	"d": func() tea.Msg {
		return NavMsg{NavRight}
	},
	"right": func() tea.Msg {
		return NavMsg{NavRight}
	},
	"enter": func() tea.Msg {
		return NavMsg{NavEnter}
	},
	"esc": func() tea.Msg {
		return NavMsg{NavBack}
	},
	"q": func() tea.Msg {
		return NavMsg{NavQuit}
	},
	"ctrl+c": func() tea.Msg {
		return NavMsg{NavQuit}
	},
}

//endregion
