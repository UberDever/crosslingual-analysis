package analysis_test

import (
	"fmt"
	"prototype2/analysis"
	"testing"
)

func TestUsecase1(t *testing.T) {

	csharpAst := analysis.Usecase1_CSharp()
	jsAst := analysis.Usecase1_JS()
	modules := analysis.Usecase1_Analyzer(csharpAst, jsAst)
	links := analysis.Link(modules)
	for _, m := range modules {
		fmt.Println(m)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}

func TestUsecase2(t *testing.T) {
	modules := analysis.Usecase2_Analyzer()
	links := analysis.Link(modules)
	for _, m := range modules {
		fmt.Println(m)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}

func TestUsecase3(t *testing.T) {
	modules := analysis.Usecase3_Analyzer()
	links := analysis.Link(modules)
	for _, m := range modules {
		fmt.Println(m)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}
