package analysis_test

import (
	"fmt"
	"prototype2/analysis"
	"testing"
)

func TestUsecase1(t *testing.T) {

	csharpAst := analysis.Usecase1_CSharp()
	jsAst := analysis.Usecase1_JS()
	nodes := analysis.Usecase1_Analyzer(csharpAst, jsAst)

	for n := range nodes {
		fmt.Println(n)
	}
}
