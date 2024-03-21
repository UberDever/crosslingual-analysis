package main

import (
	"fmt"
	"os"
	"translate/shared"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		return
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}

	fmt.Println()
}

func main() {
	Run()
}
