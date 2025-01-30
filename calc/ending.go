// ending.go
package calc

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func showOptions() {
	var choice string
	prompt := &survey.Select{
		Message: "Next?",
		Options: []string{"Quit", "Restart"},
	}
	err := survey.AskOne(prompt, &choice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch choice {
	case "Quit":
		fmt.Println("\033[1mGoodbye!")
		os.Exit(0)
	case "Restart":
		Start()
	default:
		fmt.Println("\033[31mInvalid choice! Exiting.")
	}
}