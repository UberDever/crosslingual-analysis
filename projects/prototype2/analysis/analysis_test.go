package analysis_test

import (
	"fmt"
	"prototype2/analysis"
	"prototype2/sexpr"
	"testing"
)

func TestUsecase1(t *testing.T) {

	csharpAst := analysis.Usecase1_CSharp()
	jsAst := analysis.Usecase1_JS()
	nodes := analysis.Usecase1_Analyzer(csharpAst, jsAst)

	S := sexpr.S
	fmt.Println(sexpr.PrettifySexpr(S("a", "b", "c", S("d", "e"), S("f", "g", "h")).StringReadable()))
	// fmt.Println(sexpr.PrettifySexpr(jsAst.StringReadable()))

	for n := range nodes {
		fmt.Println(n)
	}
}
