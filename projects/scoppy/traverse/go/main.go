package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the source code to compile")
		os.Exit(-1)
	}
	source := os.Args[1]

	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, "stub.go", source, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(root, func(n ast.Node) bool {
		fmt.Println(n)

		return true
	})
}
