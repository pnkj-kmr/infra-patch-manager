package cli

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

// setting up the color for terminal output
var greenText func(a ...interface{}) string
var redText func(a ...interface{}) string
var yellowText func(a ...interface{}) string

func init() {
	if runtime.GOOS == "windows" {
		greenText, redText, yellowText = fmt.Sprint, fmt.Sprint, fmt.Sprint
	} else {
		greenText = color.New(color.FgHiGreen).SprintFunc()
		redText = color.New(color.FgHiRed).SprintFunc()
		yellowText = color.New(color.FgYellow).SprintFunc()
	}
}
