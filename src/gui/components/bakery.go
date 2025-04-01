package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"tMail/src/gui/styles"
)

//region bakers

func BakeTextInput(placeholder string, charLimit int) textinput.Model {
	var ti = textinput.New()
	ti.Placeholder = placeholder
	ti.PlaceholderStyle = styles.CurrentTheme.Styles[styles.TextInput]
	ti.CharLimit = charLimit
	ti.Width = 30 //TODO: Make this less magical

	return ti
}

//endregion

//region helpers

/*
Cmdify will take any bubbletea.Msg and wrap it in a bubbletea.Cmd

The Msg type is the empty interface, the Cmd is a function which returns a Msg.
*/
func Cmdify(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

//endregion
