package shared

import (
	"encoding/json"
	"fmt"
)

type Ontology struct {
	Types TypeContext `json:"type_context"`
}

func (c *Ontology) UnmarshalJSON(data []byte) error {
	type aux Ontology
	var o aux
	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	top := c.Types.T("Top")
	for _, t := range c.Types.Ground {
		c.Types.Subtypes = append(c.Types.Subtypes, subtype{t, top})
	}
	if err = verifyContext(c.Types); err != nil {
		return err
	}
	*c = Ontology(o)
	return nil
}

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

type appTag string

const (
	TagApplication appTag = "application"
	TagGround      appTag = "ground"
)

// NOTE: This is like an union, one is active at the time
type application struct {
	Tag  appTag        `json:"tag"`
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
		if g.Tag == TagGround && *g.Name == name {
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
		Tag: TagApplication,
		App: &applicationC{
			Constructor: *ctor,
			Args:        args,
		},
	}
	if err := verifyType(t); err != nil {
		panic(err)
	}
	return t
}

func verifyType(t ground) error {
	switch t.Tag {
	case TagApplication:
		if t.App == nil {
			return fmt.Errorf("%v should be application of constructor", t)
		}
		for _, arg := range t.App.Args {
			if err := verifyType(ground(arg)); err != nil {
				return err
			}
		}
		expected := len(t.App.Constructor.ArgsVariance)
		got := len(t.App.Args)
		if got != expected {
			return fmt.Errorf("%v should have same kind as amount of arguments it applies to; expected %d, got %d", t, expected, got)
		}
	case TagGround:
		if t.Name == nil {
			return fmt.Errorf("%v should be ground type", t)
		}
	}
	return nil
}

func verifySubtype(s subtype) error {
	if err := verifyType(ground(s.Lhs)); err != nil {
		return err
	}
	if err := verifyType(ground(s.Rhs)); err != nil {
		return err
	}
	return nil
}

func verifyContext(c TypeContext) error {
	for _, t := range c.Ground {
		if err := verifyType(t); err != nil {
			return err
		}
	}
	for _, s := range c.Subtypes {
		if err := verifySubtype(s); err != nil {
			return err
		}
	}
	return nil
}
