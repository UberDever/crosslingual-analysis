package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
	ss "translate/shared"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func Mock(request ss.Arguments) {
	data, err := os.ReadFile(*request.Ontology)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ontology ss.Ontology
	err = json.Unmarshal(data, &ontology)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var counter ss.CounterService
	if request.CounterURL == nil {
		counter = ss.NewCounterServiceMock()
	} else {
		counter = ss.NewCounterServiceImpl(*request.CounterURL)
	}
	extracted := map[string]func(){
		"evaluation/Example 1/backend/server.go": func() {
			_ = counter
		},
		"evaluation/Example 5/server.go": func() {
			_ = counter
		},
	}

	for path := range extracted {
		if request.Path != nil && strings.Contains(*request.Path, path) {
			extracted[path]()
			break
		}
	}
}

// TODO: SPA 5.7 Reaching Definitions Analysis (def-use)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		os.Exit(1)
	}
	request := ss.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	path := "stub.go"
	if request.Path != nil {
		path = *request.Path
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, request.Code, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files := []*ast.File{f}

	packag := types.NewPackage("stub", "")

	pkg, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, packag, files, ssa.SanityCheckFunctions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//TODO: Rework example 1 to make analysis consistent
	for name, member := range pkg.Members {
		switch m := member.(type) {
		case *ssa.Function:
			fmt.Println(name)
			m.WriteTo(os.Stdout)
			for _, block := range m.DomPreorder() {
				for _, inst := range block.Instrs {
					// TODO: Here we get instructions with corresponding operands
					// They are sufficient for building the DFG (data-flow graph) or CFG (control flow graph)
					// I think the former is easier to get
					// So, we get the graph where: verticies - operands, edges - operations on that operands
					// I.e. t0 --os.ReadFile-> t1
					// The task is to mark some operations as potentially effectful (IO mainly)
					// and infer which nodes should be analyzed for cross-language dependencies

					switch instr := inst.(type) {
					case *ssa.MakeClosure:
						if f, ok := instr.Fn.(*ssa.Function); ok {
							f.WriteTo(os.Stdout)
						}
					}
					_ = inst
				}
			}

		}
	}
}
