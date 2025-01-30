package calc

import (
	"fmt"
	"strconv"
)

func inputFixer(theString string) float32 {
	theNumber, err := strconv.ParseFloat(theString, 32)
	if err != nil {
		fmt.Println("\033[31mInvalid input!")
		showOptions()
	}
	return float32(theNumber)
}
