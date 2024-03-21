package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	ss "translate/shared"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		os.Exit(1)
	}
	request := ss.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	source := request.Code
	fset := token.NewFileSet()
	root, err := parser.ParseFile(fset, "stub.go", source, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ast.Inspect(root, func(n ast.Node) bool {
		fmt.Println(n)

		return true
	})
}
