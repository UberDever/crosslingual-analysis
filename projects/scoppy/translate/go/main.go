package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
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
	source := request.Code
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
