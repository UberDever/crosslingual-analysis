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
	i := uint(69) // Arbitrary counter so test results can be more diverse
	counter.counter = &i

	var tm Template
	name := "declare Web server"
	for _, t := range o.Templates {
		if t.Name == name {
			tm = t
			break
		}
	}
	if tm.Name == "" {
		t.Fatalf("Template %s not found in ontology", name)
	}

	stack := Stack[any]{}
	stack.Push(NewVariable(^uint(0), BindingScope))
	stack.Push(NewIdentifier(
		"http://localhost:1234/item",
		"", 0, 0))

	cs, result, err := tm.Eval(counter, &stack)
	if err != nil {
		t.Fatal(err)
	}

	if !stack.IsEmpty() {
		t.Fatalf("Stack should be empty, but has %v", stack.Values())
	}

	expectedConstraints := `{"Usage":[{"id":69,"identifier":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"usage":"","scope":{"index":18446744073709551615,"name":"_"}},{"id":70,"identifier":{"name":"GET","path":"","start":0,"length":0},"usage":"","scope":{"index":71,"name":"_"}},{"id":72,"identifier":{"name":"application/json","path":"","start":0,"length":0},"usage":"","scope":{"index":69,"name":"_"}}],"DirectEdge":null,"AssociationKnown":[{"id":73,"declaration":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"scope":{"index":71,"name":"_"}},{"id":74,"declaration":{"name":"GET","path":"","start":0,"length":0},"scope":{"index":69,"name":"_"}},{"id":75,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":73,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	expectedResult := `{"Usage":null,"DirectEdge":null,"AssociationKnown":[{"id":75,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":73,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	if err := CompareJsonOutput(expectedConstraints, cs.String()); err != nil {
		t.Fatal(err)
	}
	if err := CompareJsonOutput(expectedResult, result.String()); err != nil {
		t.Fatal(err)
	}
}
