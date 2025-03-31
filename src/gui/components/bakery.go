package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	"tMail/src/gui/styles"
)

func BakeTextInput(placeholder string, charLimit int) textinput.Model {
	var ti = textinput.New()
	ti.Placeholder = placeholder
	ti.PlaceholderStyle = styles.CurrentTheme.Styles[styles.TextInput]
	ti.CharLimit = charLimit
	ti.Width = 30 //TODO: Make this less magical

	return ti
}
