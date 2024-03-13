package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"translate/shared"
)

func Run() {
	if len(os.Args) < 2 {
		log.Print("No argument were provided to translator")
		os.Exit(1)
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	source := request.Code
	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, "stub.go", source, parser.ParseComments)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	ast.Inspect(root, func(n ast.Node) bool {
		log.Print(n)

		return true
	})
}
