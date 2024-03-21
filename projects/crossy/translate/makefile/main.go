package main

import (
	"fmt"
	"os"
	ss "translate/shared"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		return
	}
	request := ss.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}

	fmt.Println()
}

func main() {
	Run()
}
