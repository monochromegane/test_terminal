package main

import (
	"fmt"
	"github.com/monochromegane/terminal"
	"os"
)

func main() {

	if terminal.IsTerminal(os.Stdin) {
		fmt.Println("Stdin is from Terminal.")
	} else {
		fmt.Println("Stdin isn't from Terminal.")
	}

	if terminal.IsTerminal(os.Stdout) {
		fmt.Println("Stdout is from Terminal.")
	} else {
		fmt.Println("Stdout isn't from Terminal.")
	}
}
