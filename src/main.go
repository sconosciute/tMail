package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"tMail/src/gui"
	"tMail/src/gui/styles"
)

func main() {
	styles.InitDefaultTheme()
	Controller := gui.NewGui()
	file, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("Fatal error: ", err)
		os.Exit(-1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close log file, it may be corrupted:\n", err)
		}
	}(file)

	tMail := tea.NewProgram(*Controller, tea.WithAltScreen())
	if _, err := tMail.Run(); err != nil {
		log.Fatal(err)
	}
}
