package shared

import (
	"encoding/json"
	"reflect"
	"testing"
)

func getConstraints(id uint) Constraints {
	return Constraints{
		Usage: []Usage{NewUsage(id,
			NewIdentifier("a", "/some/path", 1, 1),
			UsageDecl,
			NewVariable(1, BindingScope),
		)},
		Resolution: []Resolution{NewResolution(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingDelta),
		)},
		Uniqueness: []Uniqueness{NewUniqueness(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
		)},
		TypeDeclKnown: []TypeDeclKnown{NewTypeDeclKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingTau),
		)},
		TypeDeclUnknown: []TypeDeclUnknown{NewTypeDeclUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(1, BindingTau),
		)},
		DirectEdge: []DirectEdge{NewDirectEdge(id,
			NewVariable(1, BindingScope),
			NewVariable(2, BindingScope),
			"parent",
		)},
		AssociationKnown: []AssociationKnown{NewAssociationKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(2, BindingScope),
		)},
		NominalEdge: []NominalEdge{NewNominalEdge(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
			"import",
		)},
		Subset: []Subset{NewSubset(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
			NewNamesCollection(NamesReferences, NewVariable(2, BindingScope)),
		)},
		AssociationUnknown: []AssociationUnknown{NewAssociationUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(2, BindingScope),
		)},
		EqualKnown: []EqualKnown{NewEqualKnown(id,
			NewVariable(1, BindingTau),
			ground{
				Tag:  TagGround,
				Name: new(string),
			},
		)},
		EqualUnknown: []EqualUnknown{NewEqualUnknown(id,
			NewVariable(1, BindingTau),
			NewVariable(2, BindingTau),
		)},
		MustResolve: []MustResolve{NewMustResolve(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
		)},
		Essential: []Essential{NewEssential(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
		)},
		Exclusive: []Exclusive{NewExclusive(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
		)},
		Iconic: []Iconic{NewIconic(id,
			NewIdentifier("a", "/some/path", 1, 1),
		)},
	}
}

func TestJsonDumpRegression(t *testing.T) {
	expected := `{"Usage":[{"id":1,"identifier":{"name":"a","path":"/some/path","start":1,"length":1},"usage":"declaration","scope":{"index":1,"name":"_"}}],"DirectEdge":[{"id":1,"lhs":{"index":1,"name":"_"},"rhs":{"index":2,"name":"_"},"label":"parent"}],"AssociationKnown":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":2,"name":"_"}}],"NominalEdge":[{"id":1,"scope":{"index":1,"name":"_"},"reference":{"name":"a","path":"/some/path","start":1,"length":1},"label":"import"}],"Resolution":[{"id":1,"reference":{"name":"a","path":"/some/path","start":1,"length":1},"declaration":{"index":1,"name":"delta"}}],"Uniqueness":[{"id":1,"names":{"collection":"referenced","scope":{"index":1,"name":"_"}}}],"Subset":[{"id":1,"lhs":{"collection":"referenced","scope":{"index":1,"name":"_"}},"rhs":{"collection":"referenced","scope":{"index":2,"name":"_"}}}],"AssociationUnknown":[{"id":1,"declaration":{"index":1,"name":"delta"},"scope":{"index":2,"name":"_"}}],"TypeDeclKnown":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"variable":{"index":1,"name":"tau"}}],"TypeDeclUnknown":[{"id":1,"declaration":{"index":1,"name":"delta"},"variable":{"index":1,"name":"tau"}}],"EqualKnown":[{"id":1,"t1":{"index":1,"name":"tau"},"t2":{"tag":"ground","name":""}}],"EqualUnknown":[{"id":1,"t1":{"index":1,"name":"tau"},"t2":{"index":2,"name":"tau"}}],"MustResolve":[{"id":1,"reference":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":1,"name":"_"}}],"Essential":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":1,"name":"_"}}],"Exclusive":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":1,"name":"_"}}],"Iconic":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1}}]}`
	var id uint = 1
	cs := getConstraints(id)
	j, err := json.Marshal(cs)
	if err != nil {
		t.Fatal(err)
	}
	got := string(j)
	if err := CompareJsonOutput(expected, got); err != nil {
		t.Fatal(err)
	}
}

func TestJsonReadRegression(t *testing.T) {
	j := `{"Usage":[{"id":1,"identifier":{"name":"a","path":"/some/path","start":1,"length":1},"usage":"declaration","scope":{"index":1,"name":"_"}}],"DirectEdge":[{"id":1,"lhs":{"index":1,"name":"_"},"rhs":{"index":2,"name":"_"},"label":"parent"}],"AssociationKnown":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":2,"name":"_"}}],"NominalEdge":[{"id":1,"scope":{"index":1,"name":"_"},"reference":{"name":"a","path":"/some/path","start":1,"length":1},"label":"import"}],"Resolution":[{"id":1,"reference":{"name":"a","path":"/some/path","start":1,"length":1},"declaration":{"index":1,"name":"delta"}}],"Uniqueness":[{"id":1,"names":{"collection":"referenced","scope":{"index":1,"name":"_"}}}],"Subset":[{"id":1,"lhs":{"collection":"referenced","scope":{"index":1,"name":"_"}},"rhs":{"collection":"referenced","scope":{"index":2,"name":"_"}}}],"AssociationUnknown":[{"id":1,"declaration":{"index":1,"name":"delta"},"scope":{"index":2,"name":"_"}}],"TypeDeclKnown":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"variable":{"index":1,"name":"tau"}}],"TypeDeclUnknown":[{"id":1,"declaration":{"index":1,"name":"delta"},"variable":{"index":1,"name":"tau"}}],"EqualKnown":[{"id":1,"t1":{"index":1,"name":"tau"},"t2":{"tag":"ground","name":""}}],"EqualUnknown":[{"id":1,"t1":{"index":1,"name":"tau"},"t2":{"index":2,"name":"tau"}}],"MustResolve":[{"id":1,"reference":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":1,"name":"_"}}],"Essential":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":1,"name":"_"}}],"Exclusive":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1},"scope":{"index":1,"name":"_"}}],"Iconic":[{"id":1,"declaration":{"name":"a","path":"/some/path","start":1,"length":1}}]}`
	var cs Constraints
	json.Unmarshal([]byte(j), &cs)

	var id uint = 1
	expected := getConstraints(id)
	if !reflect.DeepEqual(cs, expected) {
		lhs, _ := json.Marshal(expected)
		rhs, _ := json.Marshal(cs)
		t.Fatal(CompareJsonOutput(string(lhs), string(rhs)))
	}
}
