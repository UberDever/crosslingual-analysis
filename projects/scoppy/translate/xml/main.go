package main

import (
	"fmt"
	"log"
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
		return
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}

	root, err := xmltree.Parse([]byte(request.Code))
	if err != nil {
		log.Fatal(err)
	}

	traverse(*root, func(v xmltree.Element) bool {
		fmt.Println(string(v.Content))
		return true
	})
}
