// core.go
package calc

import (
	"fmt"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

func Start() {
	var selected string

	prompt := &survey.Select{
		Message: "\033[30mChoose an operation:",
		Options: calcOperations,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var input1, input2 string
	fmt.Print("1st number: "); fmt.Scan(&input1); no1 := inputFixer(input1);
	
	fmt.Print("2st number: "); fmt.Scan(&input2); no2 := inputFixer(input2);

	var result float32

	switch selected {
	case "Addition":
		result = no1 + no2
	case "Subtraction":
		result = no1 - no2
	case "Multiplication":
		result = no1 * no2
	case "Division":
		if no2 == 0 {
			fmt.Println("\033[31mError: Division by zero!")
			return
		}
		result = no1 / no2
	}

	fmt.Println("\033[1mresult: \033[32m", result)

	time.Sleep(800 * time.Millisecond)
	showOptions()
}
