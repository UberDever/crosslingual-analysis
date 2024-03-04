package main

import (
	"fmt"
	"log"
	"os"

	"aqwari.net/xml/xmltree"
)

func traverse(v xmltree.Element, f func(v xmltree.Element) bool) {
	f(v)
	for i := range v.Children {
		traverse(v.Children[i], f)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the source code to compile")
		os.Exit(-1)
	}
	source := os.Args[1]

	root, err := xmltree.Parse([]byte(source))
	if err != nil {
		log.Fatal(err)
	}

	traverse(*root, func(v xmltree.Element) bool {
		fmt.Println(string(v.Content))
		return true
	})
}
