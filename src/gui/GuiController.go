package gui

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"tMail/src/globals"
	"tMail/src/gui/components"
)

type Controller struct {
	views         []tea.Model
	currentView   tea.Model
	currentKeyMap map[string]tea.Cmd
}

func NewGui() *Controller {
	return &Controller{
		currentView:   components.NewURLChecker(),
		currentKeyMap: globals.DefaultKeyBinds,
	}
}

func (c Controller) Init() tea.Cmd {
	log.Print("Initializing GUI")
	return nil
}

func (c Controller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var fwd = msg
	switch msg.(type) {
	case tea.KeyMsg:
		msgStr := msg.(tea.KeyMsg).String()
		if nav, ok := c.currentKeyMap[msgStr]; ok {
			log.Printf("Intercepted Key Message: %s\n	Converted to: %s", msgStr, nav())
			if nav().(globals.NavMsg).String() == globals.NavQuit {
				return c, tea.Quit
			} else {
				fwd = nav()
			}
		}
	}
	c.currentView, cmd = c.currentView.Update(fwd)
	return c, cmd
}

func (c Controller) View() string {
	return c.currentView.View()
}
