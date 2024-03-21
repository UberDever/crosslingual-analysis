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
	Usage              []Usage
	Resolution         []Resolution
	Uniqueness         []Uniqueness
	TypeDeclKnown      []TypeDeclKnown
	TypeDeclUnknown    []TypeDeclUnknown
	DirectEdge         []DirectEdge
	AssociationKnown   []AssociationKnown
	NominalEdge        []NominalEdge
	Subset             []Subset
	AssociationUnknown []AssociationUnknown
}

func (c Constraints) Merge(cs Constraints) Constraints {
	return Constraints{
		Usage:              append(c.Usage, cs.Usage...),
		Resolution:         append(c.Resolution, cs.Resolution...),
		Uniqueness:         append(c.Uniqueness, cs.Uniqueness...),
		TypeDeclKnown:      append(c.TypeDeclKnown, cs.TypeDeclKnown...),
		TypeDeclUnknown:    append(c.TypeDeclUnknown, cs.TypeDeclUnknown...),
		DirectEdge:         append(c.DirectEdge, cs.DirectEdge...),
		AssociationKnown:   append(c.AssociationKnown, cs.AssociationKnown...),
		NominalEdge:        append(c.NominalEdge, cs.NominalEdge...),
		Subset:             append(c.Subset, cs.Subset...),
		AssociationUnknown: append(c.AssociationUnknown, cs.AssociationUnknown...),
	}
}

type Identifier struct {
	Name string `json:"name"`
	source
}

func NewIdentifier(name string, path string, start, length uint) Identifier {
	return Identifier{
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
type Usage struct {
	distinct
	Identifier Identifier `json:"identifier"`
	UsageType  usageType  `json:"usage"`
	Scope      variable   `json:"scope"`
}

func NewUsage(id uint, identifier Identifier, usageType usageType, scope variable) Usage {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	if !(usageType == UsageDecl || usageType == UsageRef) {
		panic(fmt.Sprintf("The usage %v has unexpected usage type %v", id, usageType))
	}
	return Usage{distinct{id}, identifier, usageType, scope}
}

// Resolution constraint (3) (only for the case when declaration is variable)
type Resolution struct {
	distinct
	Reference   Identifier `json:"reference"`
	Declaration variable   `json:"declaration"`
}

func NewResolution(id uint, identifier Identifier, declaration variable) Resolution {
	if declaration.Name != BindingDelta {
		panic(fmt.Sprintf("In resolution %v, %v is not a declaration", id, declaration))
	}
	return Resolution{distinct{id}, identifier, declaration}
}

// Uniqueness constraint (4)
type Uniqueness struct {
	distinct
	Names names `json:"names"`
}

func NewUniqueness(id uint, names names) Uniqueness {
	return Uniqueness{distinct{id}, names}
}

// Type declaration constraint (6), where D is known
type TypeDeclKnown struct {
	distinct
	Declaration Identifier `json:"declaration"`
	Type        variable   `json:"variable"`
}

func NewTypeDeclKnown(id uint, identifier Identifier, typevar variable) TypeDeclKnown {
	if typevar.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, typevar))
	}
	return TypeDeclKnown{distinct{id}, identifier, typevar}
}

// Type declaration constraint (6), where D is decl variable
type TypeDeclUnknown struct {
	distinct
	Declaration variable `json:"declaration"`
	Type        variable `json:"variable"`
}

func NewTypeDeclUnknown(id uint, identifier variable, typevar variable) TypeDeclUnknown {
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In type declaration %v, %v is not a declaration", id, identifier))
	}
	if typevar.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, typevar))
	}
	return TypeDeclUnknown{distinct{id}, identifier, typevar}
}

type EqualKnown struct {
	distinct
	T1 variable `json:"tau"`
	//TODO: type...
}

// Direct edge constraint (8)
type DirectEdge struct {
	distinct
	Lhs   variable `json:"lhs"`
	Rhs   variable `json:"rhs"`
	Label string   `json:"label"`
}

func NewDirectEdge(id uint, lhs, rhs variable, label string) DirectEdge {
	if !(lhs.Name == BindingScope || lhs.Name == BindingSigma) {
		panic(fmt.Sprintf("In direct edge %v, %v is not a scope variable", id, lhs))
	}
	if !(rhs.Name == BindingScope || rhs.Name == BindingSigma) {
		panic(fmt.Sprintf("In direct edge %v, %v is not a scope variable", id, rhs))
	}
	return DirectEdge{distinct{id}, lhs, rhs, label}
}

// Association constraint (9) (only for the case when declaration is known)
type AssociationKnown struct {
	distinct
	Declaration Identifier `json:"declaration"`
	Scope       variable   `json:"scope"`
}

func NewAssociationKnown(id uint, identifier Identifier, scope variable) AssociationKnown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	return AssociationKnown{distinct{id}, identifier, scope}
}

// Nominal edge constraint (10)
type NominalEdge struct {
	distinct
	Scope     variable   `json:"scope"`
	Reference Identifier `json:"reference"`
	Label     string     `json:"label"`
}

func NewNominalEdge(id uint, reference Identifier, scope variable, label string) NominalEdge {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	return NominalEdge{distinct{id}, scope, reference, label}
}

// Subset constraint (13)
type Subset struct {
	distinct
	Lhs names `json:"lhs"`
	Rhs names `json:"rhs"`
}

func NewSubset(id uint, lhs, rhs names) Subset {
	if lhs.NamesType != rhs.NamesType {
		panic(fmt.Sprintf("In subset %v, %v and %v has different types", id, lhs, rhs))
	}
	return Subset{distinct{id}, lhs, rhs}
}

// Association constraint (14) (only for the case when declaration is a variable)
type AssociationUnknown struct {
	distinct
	Declaration variable `json:"declaration"`
	Scope       variable `json:"scope"`
}

func NewAssociationUnknown(id uint, identifier variable, scope variable) AssociationUnknown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In type declaration %v, %v is not a declaration", id, identifier))
	}
	return AssociationUnknown{distinct{id}, identifier, scope}
}
