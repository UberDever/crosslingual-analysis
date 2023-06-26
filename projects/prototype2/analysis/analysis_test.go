package analysis_test

import (
	"fmt"
	"prototype2/analysis"
	"testing"
)

func TestUsecase1(t *testing.T) {

	csharpAst := analysis.Usecase1_CSharp()
	jsAst := analysis.Usecase1_JS()
	fragments := analysis.Usecase1_Analyzer(csharpAst, jsAst)
	links := analysis.Link(fragments)
	for _, m := range fragments {
		fmt.Println(m)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}

func TestUsecase2(t *testing.T) {
	fragments := analysis.Usecase2_Analyzer()
	links := analysis.Link(fragments)
	for _, m := range fragments {
		fmt.Println(m)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}

func TestUsecase3(t *testing.T) {
	fragments := analysis.Usecase3_Analyzer()
	links := analysis.Link(fragments)
	for _, m := range fragments {
		fmt.Println(m)
	}
	for _, l := range links {
		fmt.Println(l)
	}
}
