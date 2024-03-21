package shared

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// NOTE: Golang type system is rigid and isn't designed for such
// kind of stuff, so bear with this implementation

type locationRange struct {
	Start  uint `json:"start"`
	Length uint `json:"length"`
}

type source struct {
	Path string `json:"path"`
	locationRange
}

type Constraint interface {
	Id() uint
	variantConstraint()
}

type typedConstraint struct {
	Constraint string          `json:"constraint"`
	Body       json.RawMessage `json:"body"`
}

func AssignType(c Constraint) typedConstraint {
	body, err := json.Marshal(c)
	if err != nil {
		panic(fmt.Sprintf("Unreachable %v", err))
	}
	return typedConstraint{
		Constraint: reflect.TypeOf(c).Name(),
		Body:       body,
	}
}

func EraseType(c typedConstraint) Constraint {
	if c.Constraint == reflect.TypeOf(usage{}).Name() {
		var result usage
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(resolution{}).Name() {
		var result resolution
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(uniqueness{}).Name() {
		var result uniqueness
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(typeDeclKnown{}).Name() {
		var result typeDeclKnown
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(typeDeclUnknown{}).Name() {
		var result typeDeclUnknown
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(directEdge{}).Name() {
		var result directEdge
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(associationKnown{}).Name() {
		var result associationKnown
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(associationUnknown{}).Name() {
		var result associationUnknown
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(nominalEdge{}).Name() {
		var result nominalEdge
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else if c.Constraint == reflect.TypeOf(subset{}).Name() {
		var result subset
		err := json.Unmarshal(c.Body, &result)
		if err != nil {
			panic(err)
		}
		return result
	} else {
		panic("Unreachable")
	}

}

type identifier struct {
	Name string `json:"name"`
	source
}

func NewIdentifier(name string, path string, start, length uint) identifier {
	return identifier{
		Name: name,
		source: source{
			Path: path,
			locationRange: locationRange{
				Start:  start,
				Length: length,
			},
		},
	}
}

type bindingType string

const (
	BindingDelta bindingType = "delta" // Delta is for unresolved declarations
	BindingSigma bindingType = "sigma" // Sigma is for unresovled scopes
	BindingTau   bindingType = "tau"   // Tau is for type variable
	BindingScope bindingType = "_"     // Alpha is for common ground scopes
)

type variable struct {
	Index uint        `json:"index"`
	Name  bindingType `json:"name"`
}

func NewVariable(index uint, Type bindingType) variable {
	return variable{index, Type}
}

type distinct struct {
	I uint `json:"id"`
}

func (i distinct) Id() uint { return i.I }

type namesType string

const (
	NamesDeclarations namesType = "declared"
	NamesReferences   namesType = "referenced"
	NamesVisible      namesType = "visible"
)

// Name collections (5, 11, 12)
type names struct {
	NamesType namesType `json:"collection"`
	Scope     variable  `json:"scope"`
}

func NewNamesCollection(t namesType, scope variable) names {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In names collection %v, %v is not a scope variable", t, scope))
	}
	return names{t, scope}
}

type usageType string

const (
	UsageDecl usageType = "declaration"
	UsageRef  usageType = "reference"
)

// Declaration constraint (1) (only for the case when declaration is non-variable)
// Reference constraint (2) (only for the case when reference is non-variable)
type usage struct {
	distinct
	Identifier identifier `json:"identifier"`
	UsageType  usageType  `json:"usage"`
	Scope      variable   `json:"scope"`
}

func (usage) variantConstraint() {}

func NewUsage(id uint, identifier identifier, usageType usageType, scope variable) usage {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	if !(usageType == UsageDecl || usageType == UsageRef) {
		panic(fmt.Sprintf("The usage %v has unexpected usage type %v", id, usageType))
	}
	return usage{distinct{id}, identifier, usageType, scope}
}

// Resolution constraint (3) (only for the case when declaration is variable)
type resolution struct {
	distinct
	Reference   identifier `json:"reference"`
	Declaration variable   `json:"declaration"`
}

func (resolution) variantConstraint() {}

func NewResolution(id uint, identifier identifier, declaration variable) resolution {
	if declaration.Name != BindingDelta {
		panic(fmt.Sprintf("In resolution %v, %v is not a declaration", id, declaration))
	}
	return resolution{distinct{id}, identifier, declaration}
}

// Uniqueness constraint (4)
type uniqueness struct {
	distinct
	Names names `json:"names"`
}

func (uniqueness) variantConstraint() {}

func NewUniqueness(id uint, names names) uniqueness {
	return uniqueness{distinct{id}, names}
}

// Type declaration constraint (6), where D is known
type typeDeclKnown struct {
	distinct
	Declaration identifier `json:"declaration"`
	Type        variable   `json:"variable"`
}

func (typeDeclKnown) variantConstraint() {}

func NewTypeDeclKnown(id uint, identifier identifier, typevar variable) typeDeclKnown {
	if typevar.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, typevar))
	}
	return typeDeclKnown{distinct{id}, identifier, typevar}
}

// Type declaration constraint (6), where D is decl variable
type typeDeclUnknown struct {
	distinct
	Declaration variable `json:"declaration"`
	Type        variable `json:"variable"`
}

func (typeDeclUnknown) variantConstraint() {}

func NewTypeDeclUnknown(id uint, identifier variable, typevar variable) typeDeclUnknown {
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In type declaration %v, %v is not a declaration", id, identifier))
	}
	if typevar.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, typevar))
	}
	return typeDeclUnknown{distinct{id}, identifier, typevar}
}

// Direct edge constraint (8)
type directEdge struct {
	distinct
	Lhs   variable `json:"lhs"`
	Rhs   variable `json:"rhs"`
	Label string   `json:"label"`
}

func (directEdge) variantConstraint() {}

func NewDirectEdge(id uint, lhs, rhs variable, label string) directEdge {
	if !(lhs.Name == BindingScope || lhs.Name == BindingSigma) {
		panic(fmt.Sprintf("In direct edge %v, %v is not a scope variable", id, lhs))
	}
	if !(rhs.Name == BindingScope || rhs.Name == BindingSigma) {
		panic(fmt.Sprintf("In direct edge %v, %v is not a scope variable", id, rhs))
	}
	return directEdge{distinct{id}, lhs, rhs, label}
}

// Association constraint (9) (only for the case when declaration is known)
type associationKnown struct {
	distinct
	Declaration identifier `json:"declaration"`
	Scope       variable   `json:"scope"`
}

func (associationKnown) variantConstraint() {}

func NewAssociationKnown(id uint, identifier identifier, scope variable) associationKnown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	return associationKnown{distinct{id}, identifier, scope}
}

// Nominal edge constraint (10)
type nominalEdge struct {
	distinct
	Scope     variable   `json:"scope"`
	Reference identifier `json:"reference"`
	Label     string     `json:"label"`
}

func (nominalEdge) variantConstraint() {}

func NewNominalEdge(id uint, reference identifier, scope variable, label string) nominalEdge {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	return nominalEdge{distinct{id}, scope, reference, label}
}

// Subset constraint (13)
type subset struct {
	distinct
	Lhs names `json:"lhs"`
	Rhs names `json:"rhs"`
}

func (subset) variantConstraint() {}

func NewSubset(id uint, lhs, rhs names) subset {
	if lhs.NamesType != rhs.NamesType {
		panic(fmt.Sprintf("In subset %v, %v and %v has different types", id, lhs, rhs))
	}
	return subset{distinct{id}, lhs, rhs}
}

// Association constraint (9) (only for the case when declaration is a variable)
type associationUnknown struct {
	distinct
	Declaration variable `json:"declaration"`
	Scope       variable `json:"scope"`
}

func (associationUnknown) variantConstraint() {}

func NewAssociationUnknown(id uint, identifier variable, scope variable) associationUnknown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In type declaration %v, %v is not a declaration", id, identifier))
	}
	return associationUnknown{distinct{id}, identifier, scope}
}
