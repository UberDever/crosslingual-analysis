package shared

import (
	"fmt"
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

// Since support for sum types in golang is nonexistent,
// I forced to use SOA here
type Constraints struct {
	Usage              []usage
	Resolution         []resolution
	Uniqueness         []uniqueness
	TypeDeclKnown      []typeDeclKnown
	TypeDeclUnknown    []typeDeclUnknown
	DirectEdge         []directEdge
	AssociationKnown   []associationKnown
	NominalEdge        []nominalEdge
	Subset             []subset
	AssociationUnknown []associationUnknown
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

func NewUniqueness(id uint, names names) uniqueness {
	return uniqueness{distinct{id}, names}
}

// Type declaration constraint (6), where D is known
type typeDeclKnown struct {
	distinct
	Declaration identifier `json:"declaration"`
	Type        variable   `json:"variable"`
}

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

func NewAssociationUnknown(id uint, identifier variable, scope variable) associationUnknown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In type declaration %v, %v is not a declaration", id, identifier))
	}
	return associationUnknown{distinct{id}, identifier, scope}
}
