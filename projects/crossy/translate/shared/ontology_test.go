package shared

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

const ONTOLOGY_PATH = ANCHOR_PATH + "evaluation/ontology.json"

func setup() (o Ontology, counter CounterServiceMock, err error) {
	data, err := os.ReadFile(ONTOLOGY_PATH)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &o)
	if err != nil {
		return
	}

	counter = NewCounterServiceMock()
	i := uint(69) // Arbitrary counter so test results can be more diverse
	counter.counter = &i
	return
}

func checkEval(tm template, counter CounterService, stack Stack[any], expectedConstraints, expectedResult string) error {
	before := tm
	cs, result, err := tm.Eval(counter, &stack)
	// return fmt.Errorf("%v\n%v", cs, result)
	if err != nil {
		return err
	}
	after := tm
	if !reflect.DeepEqual(before, after) {
		return fmt.Errorf("Template %v should not be side-effected by eval", tm)
	}

	if !stack.IsEmpty() {
		return fmt.Errorf("Stack should be empty, but has %v", stack.Values())
	}

	if err := CompareJsonOutput(expectedConstraints, cs.String()); err != nil {
		return err
	}
	if err := CompareJsonOutput(expectedResult, result.String()); err != nil {
		return err
	}

	return nil
}

func TestTemplateEvaluationRegression1(t *testing.T) {
	o, counter, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	tm, err := o.FindTemplate("declare WebServer")
	if err != nil {
		t.Fatal(err)
	}

	stack := Stack[any]{}
	stack.Push(NewVariable(^uint(0), BindingScope))
	stack.Push(NewIdentifier(
		"http://localhost:1234/item",
		"", 0, 0))

	expectedConstraints := `{"Usage":[{"id":72,"identifier":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"usage":"declaration","scope":{"index":18446744073709551615,"name":"_"}},{"id":74,"identifier":{"name":"GET","path":"","start":0,"length":0},"usage":"declaration","scope":{"index":70,"name":"_"}},{"id":75,"identifier":{"name":"application/json","path":"","start":0,"length":0},"usage":"declaration","scope":{"index":72,"name":"_"}}],"DirectEdge":null,"AssociationKnown":[{"id":69,"declaration":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"scope":{"index":70,"name":"_"}},{"id":71,"declaration":{"name":"GET","path":"","start":0,"length":0},"scope":{"index":72,"name":"_"}},{"id":73,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":69,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	expectedResult := `{"Usage":null,"DirectEdge":null,"AssociationKnown":[{"id":73,"declaration":{"name":"application/json","path":"","start":0,"length":0},"scope":{"index":69,"name":"_"}}],"NominalEdge":null,"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	if err = checkEval(tm, counter, stack, expectedConstraints, expectedResult); err != nil {
		t.Fatal(err)
	}
}

func TestTemplateEvaluationRegression2(t *testing.T) {
	o, counter, err := setup()
	if err != nil {
		t.Fatal(err)
	}

	tm, err := o.FindTemplate("reference WebServer")
	if err != nil {
		t.Fatal(err)
	}

	stack := Stack[any]{}
	stack.Push(NewVariable(^uint(0), BindingScope))
	stack.Push(NewIdentifier(
		"http://localhost:1234/item",
		"", 0, 0))

	expectedCs := `{"Usage":[{"id":70,"identifier":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"usage":"reference","scope":{"index":18446744073709551615,"name":"_"}}],"DirectEdge":[{"id":69,"lhs":{"index":70,"name":"_"},"rhs":{"index":18446744073709551615,"name":"_"},"label":"parent"}],"AssociationKnown":null,"NominalEdge":[{"id":71,"scope":{"index":70,"name":"_"},"reference":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"label":"import"}],"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	expectedR := `{"Usage":null,"DirectEdge":null,"AssociationKnown":null,"NominalEdge":[{"id":71,"scope":{"index":70,"name":"_"},"reference":{"name":"http://localhost:1234/item","path":"","start":0,"length":0},"label":"import"}],"Resolution":null,"Uniqueness":null,"Subset":null,"AssociationUnknown":null,"TypeDeclKnown":null,"TypeDeclUnknown":null,"EqualKnown":null,"EqualUnknown":null,"MustResolve":null,"Essential":null,"Exclusive":null,"Iconic":null}`
	if err = checkEval(tm, counter, stack, expectedCs, expectedR); err != nil {
		t.Fatal(err)
	}
}
