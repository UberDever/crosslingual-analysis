package shared

import (
	"encoding/json"
	"reflect"
	"testing"
)

func AssignTypes(xs []Constraint) []typedConstraint {
	result := make([]typedConstraint, 0, len(xs))
	for _, x := range xs {
		result = append(result, AssignType(x))
	}
	return result
}

func TestJsonDumpRegression(t *testing.T) {
	expected := `[{"constraint":"usage","body":{"id":1,"identifier":{"name":"a","path":"/some/path","start":1,"length":1},"usage":"declaration","scope":{"index":1,"name":"_"}}},{"constraint":"resolution","body":{"id":1,"reference":{"name":"a","path":"/some/path","start":1,"length":1},"declaration":{"index":1,"name":"delta"}}},{"constraint":"uniqueness","body":{"id":1,"names":{"collection":"referenced","scope":{"index":1,"name":"_"}}}},{"constraint":"typeDeclKnown","body":{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"variable":{"index":1,"name":"tau"}}},{"constraint":"typeDeclUnknown","body":{"id":1,"declaration":{"index":1,"name":"delta"},"variable":{"index":1,"name":"tau"}}},{"constraint":"directEdge","body":{"id":1,"lhs":{"index":1,"name":"_"},"rhs":{"index":2,"name":"_"},"label":"parent"}},{"constraint":"associationKnown","body":{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":2,"name":"_"}}},{"constraint":"nominalEdge","body":{"id":1,"scope":{"index":1,"name":"_"},"reference":{"name":"a","path":"/some/path","start":1,"length":1},"label":"import"}},{"constraint":"subset","body":{"id":1,"lhs":{"collection":"referenced","scope":{"index":1,"name":"_"}},"rhs":{"collection":"referenced","scope":{"index":2,"name":"_"}}}},{"constraint":"associationUnknown","body":{"id":1,"declaration":{"index":1,"name":"delta"},"scope":{"index":2,"name":"_"}}}]`
	var id uint = 1
	c := AssignTypes([]Constraint{
		NewUsage(id,
			NewIdentifier("a", "/some/path", 1, 1),
			UsageDecl,
			NewVariable(1, BindingScope),
		),
		NewResolution(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingDelta),
		),
		NewUniqueness(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
		),
		NewTypeDeclKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingTau),
		),
		NewTypeDeclUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(1, BindingTau),
		),
		NewDirectEdge(id,
			NewVariable(1, BindingScope),
			NewVariable(2, BindingScope),
			"parent",
		),
		NewAssociationKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(2, BindingScope),
		),
		NewNominalEdge(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
			"import",
		),
		NewSubset(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
			NewNamesCollection(NamesReferences, NewVariable(2, BindingScope)),
		),
		NewAssociationUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(2, BindingScope),
		),
	})
	j, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	got := string(j)
	if got != expected {
		t.Fatal(CompareJsonOutput(expected, got))
	}
}

func EraseTypes(xs []typedConstraint) []Constraint {
	result := make([]Constraint, 0, len(xs))
	for _, x := range xs {
		result = append(result, EraseType(x))
	}
	return result
}

func TestJsonReadRegression(t *testing.T) {
	j := `[{"constraint":"usage","body":{"id":1,"identifier":{"name":"a","path":"/some/path","start":1,"length":1},"usage":"declaration","scope":{"index":1,"name":"_"}}},{"constraint":"resolution","body":{"id":1,"reference":{"name":"a","path":"/some/path","start":1,"length":1},"declaration":{"index":1,"name":"delta"}}},{"constraint":"uniqueness","body":{"id":1,"names":{"collection":"referenced","scope":{"index":1,"name":"_"}}}},{"constraint":"typeDeclKnown","body":{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"variable":{"index":1,"name":"tau"}}},{"constraint":"typeDeclUnknown","body":{"id":1,"declaration":{"index":1,"name":"delta"},"variable":{"index":1,"name":"tau"}}},{"constraint":"directEdge","body":{"id":1,"lhs":{"index":1,"name":"_"},"rhs":{"index":2,"name":"_"},"label":"parent"}},{"constraint":"associationKnown","body":{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":2,"name":"_"}}},{"constraint":"nominalEdge","body":{"id":1,"scope":{"index":1,"name":"_"},"reference":{"name":"a","path":"/some/path","start":1,"length":1},"label":"import"}},{"constraint":"subset","body":{"id":1,"lhs":{"collection":"referenced","scope":{"index":1,"name":"_"}},"rhs":{"collection":"referenced","scope":{"index":2,"name":"_"}}}},{"constraint":"associationUnknown","body":{"id":1,"declaration":{"index":1,"name":"delta"},"scope":{"index":2,"name":"_"}}}]`
	var untyped []typedConstraint
	json.Unmarshal([]byte(j), &untyped)

	got := EraseTypes(untyped)
	var id uint = 1
	expected := []Constraint{
		NewUsage(id,
			NewIdentifier("a", "/some/path", 1, 1),
			UsageDecl,
			NewVariable(1, BindingScope),
		),
		NewResolution(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingDelta),
		),
		NewUniqueness(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
		),
		NewTypeDeclKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingTau),
		),
		NewTypeDeclUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(1, BindingTau),
		),
		NewDirectEdge(id,
			NewVariable(1, BindingScope),
			NewVariable(2, BindingScope),
			"parent",
		),
		NewAssociationKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(2, BindingScope),
		),
		NewNominalEdge(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
			"import",
		),
		NewSubset(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
			NewNamesCollection(NamesReferences, NewVariable(2, BindingScope)),
		),
		NewAssociationUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(2, BindingScope),
		),
	}
	if !reflect.DeepEqual(got, expected) {
		lhs, _ := json.Marshal(expected)
		rhs, _ := json.Marshal(got)
		t.Fatal(CompareJsonOutput(string(lhs), string(rhs)))
	}
}
