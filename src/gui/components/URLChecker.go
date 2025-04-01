package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"net/http"
	"tMail/src/globals"
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

func NewURLChecker() URLChecker {
	var chk = URLChecker{}
	chk.input = BakeTextInput("URL", 156)
	chk.status = getInput
	chk.input.Focus()
	return chk
}

/*
Init readies this URl checker to the default state
*/
func (U URLChecker) Init() tea.Cmd {
	return textinput.Blink
}

/*
Update receives Bubble Tea messages to act on
*/
func (U URLChecker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	case globals.NavMsg:
		var nav = msg.(globals.NavMsg)
		return U, Cmdify(U.handleInput(nav))
	case *http.Response:
		U.status = responseRcvd
		U.response = msg.(*http.Response).StatusCode
	case error:
		log.Print(msg)
		U.status = getInput
		U.input = BakeTextInput("Failed, try again", 156)
		U.input.Focus()
	}

	U.input, cmd = U.input.Update(msg)

	return U, cmd
}

/*
View returns the user-friendly printable string of this object.
*/
func (U URLChecker) View() string {
	var ret string

	switch U.status {
	case responseRcvd:
		ret = fmt.Sprintf("Url: %s \nStatus: %d", U.url, U.response)
	case getInput:
		ret = "Check URL:\n " + U.input.View()
	case searching:
		ret = fmt.Sprintf("Checking %s", U.url)
	}

	return ret
}

/*
handleInput receives a NavMsg and decides what to do about it
*/
func (U *URLChecker) handleInput(msg globals.NavMsg) tea.Msg {
	if U.status == searching {
		return nil
	} //early out because we are not accepting input

	switch msg.String() {
	case globals.NavEnter:
		U.input.Blur()
		U.url = U.input.Value()
		U.status = searching
		log.Printf("Pinging %s", U.url)
		return checkServer(U.url)
	}

	return nil
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

	return res
}
