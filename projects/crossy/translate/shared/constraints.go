package shared

import "fmt"

type locationRange struct {
	Start  uint `json:"start"`
	Length uint `json:"length"`
}

type source struct {
	Path string `json:"path"`
	locationRange
}

type Constraint interface {
	variantConstraint()
}

type identifier struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	source
}

func (d identifier) variantConstraint() {}

func NewIdentifier(id uint, name string, path string, start, length uint) identifier {
	return identifier{
		Id:   id,
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

type bindingType = string

const (
	BindingDelta bindingType = "Delta" // Delta is for unresolved declarations
	BindingSigma bindingType = "Sigma" // Sigma is for unresovled scopes
	BindingTau   bindingType = "Tau"   // Tau is for type variable
	BindingScope bindingType = "_"     // Alpha is for common known scopes
)

type variable struct {
	Id   uint        `json:"id"`
	Name bindingType `json:"name"`
}

func (d variable) variantConstraint() {}

func NewVariable(id uint, Type bindingType) variable {
	return variable{id, Type}
}

type usageType = string

const (
	UsageDecl usageType = "Decl"
	UsageRef  usageType = "Ref"
)

type usage struct {
	identifier
	usageType
	Scope variable `json:"scope"`
}

func NewUsage(identifier identifier, usageType usageType, scope variable) usage {
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", identifier, scope))
	}
	if !(usageType == UsageDecl || usageType == UsageRef) {
		panic(fmt.Sprintf("The usage %v has unexpected usage type %v", identifier, usageType))
	}
	return usage{identifier, usageType, scope}
}
