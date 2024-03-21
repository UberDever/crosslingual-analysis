package shared

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonDumpRegression(t *testing.T) {
	expected := `{
		"Usage": [
		  {
			"id": 1,
			"identifier": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"usage": "declaration",
			"scope": {
			  "index": 1,
			  "name": "_"
			}
		  }
		],
		"Resolution": [
		  {
			"id": 1,
			"reference": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"declaration": {
			  "index": 1,
			  "name": "delta"
			}
		  }
		],
		"Uniqueness": [
		  {
			"id": 1,
			"names": {
			  "collection": "referenced",
			  "scope": {
				"index": 1,
				"name": "_"
			  }
			}
		  }
		],
		"TypeDeclKnown": [
		  {
			"id": 1,
			"declaration": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"variable": {
			  "index": 1,
			  "name": "tau"
			}
		  }
		],
		"TypeDeclUnknown": [
		  {
			"id": 1,
			"declaration": {
			  "index": 1,
			  "name": "delta"
			},
			"variable": {
			  "index": 1,
			  "name": "tau"
			}
		  }
		],
		"DirectEdge": [
		  {
			"id": 1,
			"lhs": {
			  "index": 1,
			  "name": "_"
			},
			"rhs": {
			  "index": 2,
			  "name": "_"
			},
			"label": "parent"
		  }
		],
		"AssociationKnown": [
		  {
			"id": 1,
			"declaration": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"scope": {
			  "index": 2,
			  "name": "_"
			}
		  }
		],
		"NominalEdge": [
		  {
			"id": 1,
			"scope": {
			  "index": 1,
			  "name": "_"
			},
			"reference": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"label": "import"
		  }
		],
		"Subset": [
		  {
			"id": 1,
			"lhs": {
			  "collection": "referenced",
			  "scope": {
				"index": 1,
				"name": "_"
			  }
			},
			"rhs": {
			  "collection": "referenced",
			  "scope": {
				"index": 2,
				"name": "_"
			  }
			}
		  }
		],
		"AssociationUnknown": [
		  {
			"id": 1,
			"declaration": {
			  "index": 1,
			  "name": "delta"
			},
			"scope": {
			  "index": 2,
			  "name": "_"
			}
		  }
		]
	  }`
	var compact bytes.Buffer
	err := json.Compact(&compact, []byte(expected))
	if err != nil {
		panic(err)
	}
	var id uint = 1
	cs := Constraints{
		Usage: []usage{NewUsage(id,
			NewIdentifier("a", "/some/path", 1, 1),
			UsageDecl,
			NewVariable(1, BindingScope),
		)},
		Resolution: []resolution{NewResolution(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingDelta),
		)},
		Uniqueness: []uniqueness{NewUniqueness(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
		)},
		TypeDeclKnown: []typeDeclKnown{NewTypeDeclKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingTau),
		)},
		TypeDeclUnknown: []typeDeclUnknown{NewTypeDeclUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(1, BindingTau),
		)},
		DirectEdge: []directEdge{NewDirectEdge(id,
			NewVariable(1, BindingScope),
			NewVariable(2, BindingScope),
			"parent",
		)},
		AssociationKnown: []associationKnown{NewAssociationKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(2, BindingScope),
		)},
		NominalEdge: []nominalEdge{NewNominalEdge(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
			"import",
		)},
		Subset: []subset{NewSubset(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
			NewNamesCollection(NamesReferences, NewVariable(2, BindingScope)),
		)},
		AssociationUnknown: []associationUnknown{NewAssociationUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(2, BindingScope),
		)},
	}
	j, err := json.Marshal(cs)
	if err != nil {
		t.Fatal(err)
	}
	got := string(j)
	if got != compact.String() {
		t.Fatal(CompareJsonOutput(expected, got))
	}
}

func TestJsonReadRegression(t *testing.T) {
	j := `{
		"Usage": [
		  {
			"id": 1,
			"identifier": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"usage": "declaration",
			"scope": {
			  "index": 1,
			  "name": "_"
			}
		  }
		],
		"Resolution": [
		  {
			"id": 1,
			"reference": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"declaration": {
			  "index": 1,
			  "name": "delta"
			}
		  }
		],
		"Uniqueness": [
		  {
			"id": 1,
			"names": {
			  "collection": "referenced",
			  "scope": {
				"index": 1,
				"name": "_"
			  }
			}
		  }
		],
		"TypeDeclKnown": [
		  {
			"id": 1,
			"declaration": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"variable": {
			  "index": 1,
			  "name": "tau"
			}
		  }
		],
		"TypeDeclUnknown": [
		  {
			"id": 1,
			"declaration": {
			  "index": 1,
			  "name": "delta"
			},
			"variable": {
			  "index": 1,
			  "name": "tau"
			}
		  }
		],
		"DirectEdge": [
		  {
			"id": 1,
			"lhs": {
			  "index": 1,
			  "name": "_"
			},
			"rhs": {
			  "index": 2,
			  "name": "_"
			},
			"label": "parent"
		  }
		],
		"AssociationKnown": [
		  {
			"id": 1,
			"declaration": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"scope": {
			  "index": 2,
			  "name": "_"
			}
		  }
		],
		"NominalEdge": [
		  {
			"id": 1,
			"scope": {
			  "index": 1,
			  "name": "_"
			},
			"reference": {
			  "name": "a",
			  "path": "/some/path",
			  "start": 1,
			  "length": 1
			},
			"label": "import"
		  }
		],
		"Subset": [
		  {
			"id": 1,
			"lhs": {
			  "collection": "referenced",
			  "scope": {
				"index": 1,
				"name": "_"
			  }
			},
			"rhs": {
			  "collection": "referenced",
			  "scope": {
				"index": 2,
				"name": "_"
			  }
			}
		  }
		],
		"AssociationUnknown": [
		  {
			"id": 1,
			"declaration": {
			  "index": 1,
			  "name": "delta"
			},
			"scope": {
			  "index": 2,
			  "name": "_"
			}
		  }
		]
	  }`
	var cs Constraints
	json.Unmarshal([]byte(j), &cs)

	var id uint = 1
	expected := Constraints{
		Usage: []usage{NewUsage(id,
			NewIdentifier("a", "/some/path", 1, 1),
			UsageDecl,
			NewVariable(1, BindingScope),
		)},
		Resolution: []resolution{NewResolution(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingDelta),
		)},
		Uniqueness: []uniqueness{NewUniqueness(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
		)},
		TypeDeclKnown: []typeDeclKnown{NewTypeDeclKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingTau),
		)},
		TypeDeclUnknown: []typeDeclUnknown{NewTypeDeclUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(1, BindingTau),
		)},
		DirectEdge: []directEdge{NewDirectEdge(id,
			NewVariable(1, BindingScope),
			NewVariable(2, BindingScope),
			"parent",
		)},
		AssociationKnown: []associationKnown{NewAssociationKnown(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(2, BindingScope),
		)},
		NominalEdge: []nominalEdge{NewNominalEdge(id,
			NewIdentifier("a", "/some/path", 1, 1),
			NewVariable(1, BindingScope),
			"import",
		)},
		Subset: []subset{NewSubset(id,
			NewNamesCollection(NamesReferences, NewVariable(1, BindingScope)),
			NewNamesCollection(NamesReferences, NewVariable(2, BindingScope)),
		)},
		AssociationUnknown: []associationUnknown{NewAssociationUnknown(id,
			NewVariable(1, BindingDelta),
			NewVariable(2, BindingScope),
		)},
	}
	if !reflect.DeepEqual(cs, expected) {
		lhs, _ := json.Marshal(expected)
		rhs, _ := json.Marshal(cs)
		t.Fatal(CompareJsonOutput(string(lhs), string(rhs)))
	}
}
