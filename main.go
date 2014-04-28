package main

import (
	"fmt"
	"github.com/monochromegane/terminal"
	"bufio"
	"os"
)

func main() {

	if terminal.IsTerminal(os.Stdin) {
		fmt.Println("Stdin is from Terminal.")
	} else {
		fmt.Println("Stdin isn't from Terminal.")
		reader := bufio.NewReader(os.Stdin);
		input,_:= reader.ReadString('\n');
		fmt.Printf("----\n%s----\n", input);
	}

	if terminal.IsTerminal(os.Stdout) {
		fmt.Println("Stdout is from Terminal.")
	} else {
		fmt.Println("Stdout isn't from Terminal.")
	}
}
