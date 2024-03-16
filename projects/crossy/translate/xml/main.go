package main

import (
	"fmt"
	"os"
	"translate/shared"

	"aqwari.net/xml/xmltree"
)

func traverse(v xmltree.Element, f func(v xmltree.Element) bool) {
	f(v)
	for i := range v.Children {
		traverse(v.Children[i], f)
	}
}

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		os.Exit(1)
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}

	root, err := xmltree.Parse([]byte(request.Code))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	traverse(*root, func(v xmltree.Element) bool {
		fmt.Println(string(v.Content))
		return true
	})
}
