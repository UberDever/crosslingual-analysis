package shared

import (
	"encoding/json"
	"fmt"
)

type variance string

const (
	VarianceCovariant     variance = "+"
	VarianceContravariant variance = "-"
	VarianceInvariant     variance = "="
)

type constructor struct {
	Name         string     `json:"name"`
	ArgsVariance []variance `json:"variance"`
}

type applicationC struct {
	Constructor constructor   `json:"constructor"`
	Args        []application `json:"args"`
}

type appKind string

const (
	appC appKind = "application"
	appG appKind = "ground"
)

// NOTE: This is like an union, one is active at the time
type application struct {
	Tag  appKind       `json:"tag"`
	App  *applicationC `json:"app,omitempty"`
	Name *string       `json:"name,omitempty"`
}

// NOTE: Type carrying is not supported, so all applied type constructors are ground (kind == *)
type ground = application

type subtype struct {
	Lhs ground `json:"lhs"`
	Rhs ground `json:"rhs"`
}

type TypeContext struct {
	Ground       []ground      `json:"ground"`
	Constructors []constructor `json:"constructors"`
	Subtypes     []subtype     `json:"subtypes"`
}

func (c TypeContext) T(name string) ground {
	var t *ground
	for _, g := range c.Ground {
		if g.Tag == appG && *g.Name == name {
			t = &g
			break
		}
	}
	if t == nil {
		panic("Unreachable " + name)
	}
	return *t
}

func (ctx TypeContext) NewT(ctorName string, args ...ground) ground {
	var ctor *constructor
	for _, c := range ctx.Constructors {
		if c.Name == ctorName {
			ctor = &c
			break
		}
	}
	if ctor == nil {
		panic("Unreachable " + ctorName)
	}
	t := application{
		Tag: appC,
		App: &applicationC{
			Constructor: *ctor,
			Args:        args,
		},
	}
	if err := VerifyType(t); err != nil {
		panic(err)
	}
	return t
}

func (ctx TypeContext) NewDeclarationConstraint(counter CounterService, decl Identifier, typ ground) Constraints {
	tau := NewVariable(counter.FreshForce(), BindingTau)
	return Constraints{
		TypeDeclKnown: []TypeDeclKnown{NewTypeDeclKnown(counter.FreshForce(), decl, tau)},
		EqualKnown:    []EqualKnown{NewEqualKnown(counter.FreshForce(), tau, typ)},
	}
}

func VerifyType(t ground) error {
	switch t.Tag {
	case appC:
		if t.App == nil {
			return fmt.Errorf("%v should be application of constructor", t)
		}
		for _, arg := range t.App.Args {
			if err := VerifyType(ground(arg)); err != nil {
				return err
			}
		}
		expected := len(t.App.Constructor.ArgsVariance)
		got := len(t.App.Args)
		if got != expected {
			return fmt.Errorf("%v should have same kind as amount of arguments it applies to; expected %d, got %d", t, expected, got)
		}
	case appG:
		if t.Name == nil {
			return fmt.Errorf("%v should be ground type", t)
		}
	}
	return nil
}

func VerifySubtype(s subtype) error {
	if err := VerifyType(ground(s.Lhs)); err != nil {
		return err
	}
	if err := VerifyType(ground(s.Rhs)); err != nil {
		return err
	}
	return nil
}

func VerifyContext(c TypeContext) error {
	for _, t := range c.Ground {
		if err := VerifyType(t); err != nil {
			return err
		}
	}
	for _, s := range c.Subtypes {
		if err := VerifySubtype(s); err != nil {
			return err
		}
	}
	return nil
}

// TODO: make an interface
func UnmarshalContext(j []byte) (TypeContext, error) {
	c := TypeContext{}
	err := json.Unmarshal(j, &c)
	if err != nil {
		return c, err
	}
	top := c.T("Top")
	for _, t := range c.Ground {
		c.Subtypes = append(c.Subtypes, subtype{t, top})
	}
	if err = VerifyContext(c); err != nil {
		return c, err
	}
	return c, nil
}
