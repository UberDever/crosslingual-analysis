package shared

import (
	"fmt"
	"reflect"
)

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

type typedConstraint struct {
	Constraint string `json:"constraint"`
	Body       any    `json:"body,inline"`
}

func (typedConstraint) variantConstraint() {}
func newConstraint(c Constraint) typedConstraint {

	return typedConstraint{
		Constraint: reflect.TypeOf(c).Name(),
		Body:       c,
	}

}
func MakeTypedConstraints(xs []Constraint) []Constraint {
	typed := make([]Constraint, 0, len(xs))
	for i := range xs {
		typed = append(typed, newConstraint(xs[i]))
	}
	return typed
}

type identifier struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	source
}

func (d identifier) variantConstraint() {}

func newIdentifier(id uint, name string, path string, start, length uint) identifier {
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
	Id        identifier `json:"identifier"`
	UsageType usageType  `json:"usage"`
	Scope     variable   `json:"scope"`
}

func (d usage) variantConstraint() {}

func NewUsage(id uint, name string, path string, start, length uint, usageType usageType, scope variable) usage {
	identifier := newIdentifier(id, name, path, start, length)
	if !(scope.Name == BindingScope || scope.Name == BindingSigma) {
		panic(fmt.Sprintf("In usage %v, %v is not a scope variable", identifier, scope))
	}
	if !(usageType == UsageDecl || usageType == UsageRef) {
		panic(fmt.Sprintf("The usage %v has unexpected usage type %v", identifier, usageType))
	}
	return usage{identifier, usageType, scope}
}
