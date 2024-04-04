package shared

import (
	"encoding/json"
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
	// scope graph constraints

	Usage            []Usage            // (1) A declaration constraint s -> xD specifies that declaration xD belongs to scope s. (2) A reference constraint xR -> s specifies that reference xR belongs to scope s
	DirectEdge       []DirectEdge       // (8) A direct edge constraint s1 -l-> s2 specifies a direct l-labeled
	AssociationKnown []AssociationKnown // (9) An association constraint xD -|> s specifies s as the associated scope of declaration xD. Associated scopes can be used to connect the declaration (e.g. a module) of a collection of names to the scope declaring those names (e.g. the body of a module).
	NominalEdge      []NominalEdge      // (10) A nominal edge constraint s -l-> xR specifies a nominal l-labeled edge from scope s to reference xR. uch an edge makes visible in s all declarations that are visible in the associated scope of the declaration to which xR resolves, according to the label on the edge

	// type-directed resolution constraints

	Resolution         []Resolution         // (3) A resolution constraint R |-> D specifies that a given reference must resolve to a given declaration. Typically, the declaration is specified as a declaration variable δ
	Uniqueness         []Uniqueness         // (4) A uniqueness constraint !N specifies that a given name collection N contains no duplicates.
	Subset             []Subset             // (13) A subset constraint N ⊂∼ N specifies that one name collection is included in another.
	AssociationUnknown []AssociationUnknown // (14) An association constraint D ~> S specifies that a given declaration has a given associated scope

	// typing constraints

	TypeDeclKnown   []TypeDeclKnown   // (6) A type declaration constraint D : T associates a type with a declaration.
	TypeDeclUnknown []TypeDeclUnknown // (6) This constraint is used in two flavors: associating a type variable (τ) with a concrete declaration, or associating a type variable with a declaration variable
	EqualKnown      []EqualKnown      // (7) A type equality constraint T ≡ T specifies that two types should be equal
	EqualUnknown    []EqualUnknown    // (7)

	// consistency constraints

	MustResolve []MustResolve // given reference must resolve to a declaration
	Essential   []Essential   // given declaration must have at least one reference (i.e. its essential for the project)
	Exclusive   []Exclusive   // given declaration must have at most one reference
	Iconic      []Iconic      // given declaration must be unique across WHOLE scope-graph, despite the scopes
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
		EqualKnown:         append(c.EqualKnown, cs.EqualKnown...),
		EqualUnknown:       append(c.EqualUnknown, cs.EqualUnknown...),
		AssociationUnknown: append(c.AssociationUnknown, cs.AssociationUnknown...),
		MustResolve:        append(c.MustResolve, cs.MustResolve...),
		Essential:          append(c.Essential, cs.Essential...),
		Exclusive:          append(c.Exclusive, cs.Exclusive...),
		Iconic:             append(c.Iconic, cs.Iconic...),
	}
}

func (c Constraints) String() string {
	s, _ := json.Marshal(c)
	return string(s)
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

type Variable struct {
	Index uint        `json:"index"`
	Name  bindingType `json:"name"`
}

func NewVariable(index uint, Type bindingType) Variable {
	return Variable{index, Type}
}

type Distinct struct {
	I uint `json:"id" mapstructure:"id"`
}

func (i Distinct) Id() uint { return i.I }

type namesType string

const (
	NamesDeclarations namesType = "declared"
	NamesReferences   namesType = "referenced"
	NamesVisible      namesType = "visible"
)

// Name collections (5, 11, 12)
type names struct {
	NamesType namesType `json:"collection"`
	Scope     Variable  `json:"scope"`
}

func NewNamesCollection(t namesType, scope Variable) names {
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
	Distinct   `mapstructure:",squash"`
	Identifier Identifier `json:"identifier"`
	UsageType  usageType  `json:"usage"`
	Scope      Variable   `json:"scope"`
}

func NewUsage(id uint, identifier Identifier, usageType usageType, scope Variable) Usage {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	if !(usageType == UsageDecl || usageType == UsageRef) {
		panic(fmt.Sprintf("The usage %v has unexpected usage type %v", id, usageType))
	}
	return Usage{Distinct{id}, identifier, usageType, scope}
}

// Resolution constraint (3) (only for the case when declaration is variable)
type Resolution struct {
	Distinct    `mapstructure:",squash"`
	Reference   Identifier `json:"reference"`
	Declaration Variable   `json:"declaration"`
}

func NewResolution(id uint, identifier Identifier, declaration Variable) Resolution {
	if declaration.Name != BindingDelta {
		panic(fmt.Sprintf("In resolution %v, %v is not a declaration", id, declaration))
	}
	return Resolution{Distinct{id}, identifier, declaration}
}

// Uniqueness constraint (4)
type Uniqueness struct {
	Distinct `mapstructure:",squash"`
	Names    names `json:"names"`
}

func NewUniqueness(id uint, names names) Uniqueness {
	return Uniqueness{Distinct{id}, names}
}

// Type declaration constraint (6), where D is known
type TypeDeclKnown struct {
	Distinct    `mapstructure:",squash"`
	Declaration Identifier `json:"declaration"`
	Type        Variable   `json:"variable"`
}

func NewTypeDeclKnown(id uint, identifier Identifier, typevar Variable) TypeDeclKnown {
	if typevar.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, typevar))
	}
	return TypeDeclKnown{Distinct{id}, identifier, typevar}
}

// Type declaration constraint (6), where D is decl variable
type TypeDeclUnknown struct {
	Distinct    `mapstructure:",squash"`
	Declaration Variable `json:"declaration"`
	Type        Variable `json:"variable"`
}

func NewTypeDeclUnknown(id uint, identifier Variable, typevar Variable) TypeDeclUnknown {
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In type declaration %v, %v is not a declaration", id, identifier))
	}
	if typevar.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, typevar))
	}
	return TypeDeclUnknown{Distinct{id}, identifier, typevar}
}

type EqualKnown struct {
	Distinct `mapstructure:",squash"`
	T1       Variable `json:"t1"`
	T2       ground   `json:"t2"`
}

func NewEqualKnown(id uint, t1 Variable, t2 ground) EqualKnown {
	if t1.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, t1))
	}
	return EqualKnown{Distinct{id}, t1, t2}
}

type EqualUnknown struct {
	Distinct `mapstructure:",squash"`
	T1       Variable `json:"t1"`
	T2       Variable `json:"t2"`
}

func NewEqualUnknown(id uint, t1 Variable, t2 Variable) EqualUnknown {
	if t1.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, t1))
	}
	if t2.Name != BindingTau {
		panic(fmt.Sprintf("In type declaration %v, %v is not a typevariable", id, t2))
	}
	return EqualUnknown{Distinct{id}, t1, t2}
}

// Direct edge constraint (8)
type DirectEdge struct {
	Distinct `mapstructure:",squash"`
	Lhs      Variable `json:"lhs"`
	Rhs      Variable `json:"rhs"`
	Label    string   `json:"label"`
}

func NewDirectEdge(id uint, lhs, rhs Variable, label string) DirectEdge {
	if !(lhs.Name == BindingScope || lhs.Name == BindingSigma) {
		panic(fmt.Sprintf("In direct edge %v, %v is not a scope variable", id, lhs))
	}
	if !(rhs.Name == BindingScope || rhs.Name == BindingSigma) {
		panic(fmt.Sprintf("In direct edge %v, %v is not a scope variable", id, rhs))
	}
	return DirectEdge{Distinct{id}, lhs, rhs, label}
}

// Association constraint (9) (only for the case when declaration is known)
type AssociationKnown struct {
	Distinct    `mapstructure:",squash"`
	Declaration Identifier `json:"declaration"`
	Scope       Variable   `json:"scope"`
}

func NewAssociationKnown(id uint, identifier Identifier, scope Variable) AssociationKnown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	return AssociationKnown{Distinct{id}, identifier, scope}
}

// Nominal edge constraint (10)
type NominalEdge struct {
	Distinct  `mapstructure:",squash"`
	Scope     Variable   `json:"scope"`
	Reference Identifier `json:"reference"`
	Label     string     `json:"label"`
}

func NewNominalEdge(id uint, reference Identifier, scope Variable, label string) NominalEdge {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", id, scope))
	}
	return NominalEdge{Distinct{id}, scope, reference, label}
}

// Subset constraint (13)
type Subset struct {
	Distinct `mapstructure:",squash"`
	Lhs      names `json:"lhs"`
	Rhs      names `json:"rhs"`
}

func NewSubset(id uint, lhs, rhs names) Subset {
	if lhs.NamesType != rhs.NamesType {
		panic(fmt.Sprintf("In subset %v, %v and %v has different types", id, lhs, rhs))
	}
	return Subset{Distinct{id}, lhs, rhs}
}

// Association constraint (14) (only for the case when declaration is a variable)
type AssociationUnknown struct {
	Distinct    `mapstructure:",squash"`
	Declaration Variable `json:"declaration"`
	Scope       Variable `json:"scope"`
}

func NewAssociationUnknown(id uint, identifier Variable, scope Variable) AssociationUnknown {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In %v, %v is not a scope variable", id, scope))
	}
	if identifier.Name != BindingDelta {
		panic(fmt.Sprintf("In %v, %v is not a declaration", id, identifier))
	}
	return AssociationUnknown{Distinct{id}, identifier, scope}
}

// Consistency constraint, given reference must resolve to a declaration
type MustResolve struct {
	Distinct  `mapstructure:",squash"`
	Reference Identifier `json:"reference"`
	Scope     Variable   `json:"scope"`
}

func NewMustResolve(id uint, reference Identifier, scope Variable) MustResolve {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In %v, %v is not a scope variable", id, scope))
	}
	return MustResolve{Distinct{id}, reference, scope}
}

// Consistency constraint, given declaration must have at least one reference (i.e. its essential for the project)
type Essential struct {
	Distinct    `mapstructure:",squash"`
	Declaration Identifier `json:"declaration"`
	Scope       Variable   `json:"scope"`
}

func NewEssential(id uint, declaration Identifier, scope Variable) Essential {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In %v, %v is not a scope variable", id, scope))
	}
	return Essential{Distinct{id}, declaration, scope}
}

// Consistency constraint, given declaration must have at most one reference
type Exclusive struct {
	Distinct    `mapstructure:",squash"`
	Declaration Identifier `json:"declaration"`
	Scope       Variable   `json:"scope"`
}

func NewExclusive(id uint, declaration Identifier, scope Variable) Exclusive {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In %v, %v is not a scope variable", id, scope))
	}
	return Exclusive{Distinct{id}, declaration, scope}
}

// Consistency constraint, given declaration must be unique across WHOLE scope-graph, despite the scopes
type Iconic struct {
	Distinct    `mapstructure:",squash"`
	Declaration Identifier `json:"declaration"`
}

func NewIconic(id uint, declaration Identifier) Iconic {
	return Iconic{Distinct{id}, declaration}
}
