package gui

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"tMail/src/gui/components"
)

type Controller struct {
	views         []Bubble
	currentView   Bubble
	currentKeyMap map[string]tea.Cmd
}

func (c Controller) Init() tea.Cmd {
	log.Print("Initializing GUI")
	return nil
}

func (c Controller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		var navCmd = c.currentKeyMap[msg.String()]()
		if m, ok := navCmd.(NavMsg); ok {
			if m.String() == navQuit {
				return c, tea.Quit
			}
		} else {
			_, cmd = c.currentView.Update(c.currentKeyMap[msg.String()])
		}
	default:
		_, cmd = c.currentView.Update(msg)
	}
	return c, cmd
}

func (c Controller) View() string {
	return c.currentView.View()
}

func New() *Controller {
	return &Controller{
		currentView:   components.URLChecker{},
		currentKeyMap: DefaultKeyBinds,
	}
}
