package shared

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
)

const ONTOLOGY_PATH = ANCHOR_PATH + "evaluation/ontology/ontology.json"

func setupCounterAndOntology() (o Ontology, counter CounterServiceMock, err error) {
	counter = NewCounterServiceMock()
	i := uint(69) // Arbitrary counter so test results can be more diverse
	counter.counter = &i

	o, err = NewOntology(counter, ONTOLOGY_PATH)
	return
}

func TestTemplateEvaluationRegression1(t *testing.T) {
	o, counter, err := setupCounterAndOntology()
	if err != nil {
		t.Fatal(err)
	}

	cs, result, err := o.EvalTemplate(
		"declare_WebServer",
		counter,
		NewVariable(^uint(0), BindingScope),
		NewIdentifier(
			"http://localhost:1234/item",
			"", 0, 0),
	)
	if err != nil {
		t.Fatal(err)
	}

	var scope variable
	mapstructure.Decode(result[0], &scope)
	data, err := json.Marshal(scope)
	if err != nil {
		t.Fatal(err)
	}
	result_str := string(data)

	expectedConstraints := `{"Usage":[{"id":72,"identifier":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"usage":"declaration","scope":{"index":18446744073709551615,"name":"_"}},{"id":73,"identifier":{"name":"GET","path":"","start":0,"length":0},"usage":"declaration","scope":{"index":70,"name":"_"}},{"id":74,"identifier":{"name":"application/json","path":"","start":0,"length":0},"usage":"declaration","scope":{"index":71,"name":"_"}}],"DirectEdge":null,"AssociationKnown":[{"id":75,"declaration":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"scope":{"index":70,"name":"_"}},{"id":76,"declaration":{"name":"GET","path":"","start":0,"length":0},"scope":{"index":71,"name":"_"}},{"id":77,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":69,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	expectedResult := `{"index":69,"name":"_"}`
	if err := CompareJsonOutput(expectedConstraints, cs.String()); err != nil {
		t.Fatal(err)
	}
	if err := CompareJsonOutput(expectedResult, result_str); err != nil {
		t.Fatal(err)
	}
}

func TestTemplateEvaluationRegression2(t *testing.T) {
	o, counter, err := setupCounterAndOntology()
	if err != nil {
		t.Fatal(err)
	}

	cs, result, err := o.EvalTemplate(
		"reference_WebServer",
		counter,
		NewVariable(^uint(0), BindingScope),
		NewIdentifier(
			"http://localhost:1234/item",
			"", 0, 0),
	)
	if err != nil {
		t.Fatal(err)
	}

	var scope variable
	mapstructure.Decode(result[0], &scope)
	data, err := json.Marshal(scope)
	if err != nil {
		t.Fatal(err)
	}
	result_str := string(data)

	expectedConstraints := `{"Usage":[{"id":70,"identifier":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"usage":"reference","scope":{"index":18446744073709551615,"name":"_"}}],"DirectEdge":[{"id":71,"lhs":{"index":69,"name":"_"},"rhs":{"index":18446744073709551615,"name":"_"},"label":"parent"}],"AssociationKnown":null,"NominalEdge":[{"id":72,"scope":{"index":69,"name":"_"},"reference":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"label":"import"}],"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	expectedResult := `{"index":69,"name":"_"}`
	if err := CompareJsonOutput(expectedConstraints, cs.String()); err != nil {
		t.Fatal(err)
	}
	if err := CompareJsonOutput(expectedResult, result_str); err != nil {
		t.Fatal(err)
	}
}

func TestTypesConcrete(t *testing.T) {
	o, counter, err := setupCounterAndOntology()
	if err != nil {
		t.Fatal(err)
	}
	_ = counter

	_, err = o.ConcreteType("String")
	if err != nil {
		t.Fatal(err)
	}
	
	_, err = o.ConcreteType("Something")
	if err == nil {
		t.Fatal("Found 'Something'")
	}
}
