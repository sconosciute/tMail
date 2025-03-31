package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"net/http"
	"time"
)

type URLChecker struct {
	url      string
	response int
	status   int
	input    textinput.Model
}

const (
	getInput = iota
	searching
	responseRcvd
)

func (U URLChecker) Init() tea.Cmd {
	U.input = BakeTextInput("URL", 156)
	U.status = getInput
	return nil
}

func (U URLChecker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Print("Updating GUI")
	return U, nil
}

func (U URLChecker) View() string {
	log.Print("Viewing GUI")
	return "URL Checker"
}

func checkServer(url string) tea.Msg {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return res.StatusCode
}
