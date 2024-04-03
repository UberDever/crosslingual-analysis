package shared

import (
	"encoding/json"
	"os"
	"testing"
)

const ONTOLOGY_PATH = ANCHOR_PATH + "evaluation/ontology.json"

func TestTemplateEvaluationRegression(t *testing.T) {
	data, err := os.ReadFile(ONTOLOGY_PATH)
	if err != nil {
		t.Fatal(err)
	}
	var o Ontology
	err = json.Unmarshal(data, &o)
	if err != nil {
		t.Fatal(err)
	}

	counter := NewCounterServiceMock()
	tm := o.Templates[0]
	stack := Stack[any]{}
	stack.Push(NewVariable(counter.FreshForce(), BindingScope))
	stack.Push(NewIdentifier(
		"http://localhost:1234/item",
		"", 0, 0))

	cs, err := tm.Evaluate(&stack)
	if err != nil {
		t.Fatal(err)
	}

	expectedConstraints := `{"Usage":[{"id":0,"identifier":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"usage":"","scope":{"index":0,"name":"_"}},{"id":0,"identifier":{"name":"GET","path":"","start":0,"length":0},"usage":"","scope":{"index":0,"name":"_"}},{"id":0,"identifier":{"name":"application/json","path":"","start":0,"length":0},"usage":"","scope":{"index":1,"name":"_"}}],"DirectEdge":null,"AssociationKnown":[{"id":0,"declaration":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"scope":{"index":0,"name":"_"}},{"id":0,"declaration":{"name":"GET","path":"","start":0,"length":0},"scope":{"index":1,"name":"_"}},{"id":0,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":2,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	expectedStack := `{"Usage":null,"DirectEdge":null,"AssociationKnown":[{"id":0,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":2,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	if err := CompareJsonOutput(expectedConstraints, cs.String()); err != nil {
		t.Fatal(err)
	}
	if err := CompareJsonOutput(expectedStack, stack.ForcePop().(Constraints).String()); err != nil {
		t.Fatal(err)
	}
}
