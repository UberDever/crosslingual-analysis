package shared

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Ontology struct {
	Types     TypeContext `json:"type_context"`
	Templates []Template  `json:"templates"`
}

func (c *Ontology) UnmarshalJSON(data []byte) error {
	type aux Ontology
	var o aux
	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}
	*c = Ontology(o)

	top := c.Types.T("Top")
	for _, t := range c.Types.Ground {
		c.Types.Subtypes = append(c.Types.Subtypes, subtype{t, top})
	}
	if err = verifyContext(c.Types); err != nil {
		return err
	}
	return nil
}

type Output struct {
	T  string `json:"type"`
	Id uint   `json:"id"`
}

type Template struct {
	Name   string         `json:"name"`
	Input  []string       `json:"input"`
	Body   map[string]any `json:"body"`
	Output []Output       `json:"output"`
}

func typecheck(input []string, stack Stack[any]) error {
	if len(stack.Values()) != len(input) {
		return fmt.Errorf("expected %d arguments to template, found %d", len(input), len(stack.Values()))
	}
	for i := range stack.Values() {
		input_type := input[i]
		v := stack.Values()[i]
		switch input_type {
		case "scope":
			scope, ok := v.(variable)
			if !ok && !(scope.Name == BindingScope || scope.Name == BindingSigma) {
				return fmt.Errorf("value %v at position %d expected to be a %s", v, i, input_type)
			}
		case "identifier":
			_, ok := v.(Identifier)
			if !ok {
				return fmt.Errorf("value %v at position %d expected to be a %s", v, i, input_type)
			}
		default:
			panic("Unreachable")
		}
	}
	return nil
}

func substitute(body map[string]any, stack Stack[any]) (map[string]any, error) {
	isParameter := func(value any) *int {
		switch v := value.(type) {
		case map[string]any:
			if len(v) != 1 {
				return nil
			}
			arg, ok := v["argument"]
			if !ok {
				return nil
			}
			argVal, ok := arg.(float64)
			if !ok {
				return nil
			}
			i := int(argVal)
			return &i
		default:
			return nil
		}
	}

	var traverse func(value any) error
	traverse = func(value any) error {
		switch v := value.(type) {
		case map[string]any:
			for key, val := range v {
				if idx := isParameter(val); idx != nil {
					// NOTE: index is always negative and starts from -1
					len := len(stack.Values())
					i := -(*idx + 1)
					if i >= len {
						return fmt.Errorf("argument %d is not in the stack, len = %d", idx, len)
					}
					v[key] = stack.Values()[i]
				}
				traverse(val)
			}
		case []any:
			for index, val := range v {
				if idx := isParameter(val); idx != nil {
					// NOTE: index is always negative and starts from -1
					len := len(stack.Values())
					i := -(*idx + 1)
					if i >= len {
						return fmt.Errorf("argument %d is not in the stack, len = %d", idx, len)
					}
					v[index] = stack.Values()[i]
				}
				traverse(val)
			}
		}
		return nil
	}

	err := traverse(body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getOutput(body map[string]any, results []Output) (out Constraints, err error) {
	csUntyped := map[string][]any{}

	var traverse func(value any) error
	traverse = func(value any) error {
		switch o := value.(type) {
		case map[string]any:
			if (Distinct{}).SameStruct(o) {
				val := uint(o["id"].(float64))
				var res Output
				for _, r := range results {
					if val == r.Id {
						res = r
						break
					}
				}
				csUntyped[res.T] = append(csUntyped[res.T], o)
			}
			for _, v := range o {
				err := traverse(v)
				if err != nil {
					return err
				}
			}
		case []any:
			for _, v := range o {
				err := traverse(v)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	err = traverse(body)
	if err != nil {
		return
	}
	mapstructure.Decode(csUntyped, &out)

	return
}

func shiftIndices(counter CounterService, cs *map[string]any, output *[]Output) error {
	seen := map[uint]uint{}

	var traverse func(value any)
	traverse = func(value any) {
		switch v := value.(type) {
		case []any:
			for i := 0; i < len(v); i++ {
				traverse(v[i])
			}
		case map[string]any:
			var newId uint
			if (Distinct{}).SameStruct(v) {
				val := uint(v["id"].(float64))
				if i, found := seen[val]; found {
					newId = i
				} else {
					newId = counter.FreshForce()
					seen[val] = newId
				}
				v["id"] = float64(newId)
			}
			if (variable{}).SameStruct(v) {
				val := uint(v["index"].(float64))
				if i, found := seen[val]; found {
					newId = i
				} else {
					newId = counter.FreshForce()
					seen[val] = newId
				}
				v["index"] = float64(newId)
			}
			for _, child := range v {
				traverse(child)
			}
		}
	}
	traverse(*cs)

	newOut := make([]Output, 0, len(*output))
	for _, out := range *output {
		if i, found := seen[out.Id]; found {
			out.Id = i
			newOut = append(newOut, out)
		} else {
			return fmt.Errorf("the result %v is not found in the body", out)
		}
	}
	*output = newOut

	return nil
}

func (tm Template) evaluate(counter CounterService, stack Stack[any]) (cs Constraints, output Constraints, err error) {
	err = typecheck(tm.Input, stack)
	if err != nil {
		return
	}
	copy := DeepCopyJSON(tm.Body).(map[string]any)
	err = shiftIndices(counter, &copy, &tm.Output)
	if err != nil {
		return
	}
	result, err := substitute(copy, stack)
	if err != nil {
		return
	}
	output, err = getOutput(result, tm.Output)
	if err != nil {
		return
	}

	err = mapstructure.Decode(result, &cs)
	if err != nil {
		return
	}

	return
}

func (tm Template) Eval(counter CounterService, stack *Stack[any]) (cs Constraints, output Constraints, err error) {
	cs, output, err = tm.evaluate(counter, *stack)
	stack.Clear()
	return
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
