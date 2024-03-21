package main

import (
	"encoding/json"
	"fmt"
	"os"
	ss "translate/shared"
)

type traverser struct {
	ctx     ss.TypeContext
	key     *ss.Identifier
	counter ss.CounterService
}

func (t traverser) value(root any) ss.Constraints {
	cs := ss.Constraints{}
	switch v := root.(type) {
	case map[string]any:
		cs = cs.Merge(t.object(v))
	case []any:
		cs = cs.Merge(t.array(v))
	case float64:
		cs = cs.Merge(t.ctx.NewDeclarationConstraint(t.counter, *t.key, t.ctx.T("Numeric")))
	case bool:
		cs = cs.Merge(t.ctx.NewDeclarationConstraint(t.counter, *t.key, t.ctx.T("Bool")))
	case string:
		cs = cs.Merge(t.ctx.NewDeclarationConstraint(t.counter, *t.key, t.ctx.T("String")))
	case nil:
		cs = cs.Merge(t.ctx.NewDeclarationConstraint(t.counter, *t.key, t.ctx.T("Top")))
	}
	return cs
}

func (t traverser) object(obj map[string]any) ss.Constraints {
	cs := ss.Constraints{}

	S := ss.NewVariable(t.counter.FreshForce(), ss.BindingScope)
	cs = cs.Merge(ss.Constraints{
		Uniqueness: []ss.Uniqueness{ss.NewUniqueness(t.counter.FreshForce(), ss.NewNamesCollection(ss.NamesDeclarations, S))},
	})
	if t.key != nil {
		cs = cs.Merge(ss.Constraints{
			AssociationKnown: []ss.AssociationKnown{ss.NewAssociationKnown(
				t.counter.FreshForce(), *t.key, S,
			)},
		})
	}
	for k, v := range obj {
		D := ss.NewIdentifier(k, "?", 0, 0) //TODO: Add path info
		t.key = &D
		cs = cs.Merge(ss.Constraints{
			Usage: []ss.Usage{
				ss.NewUsage(t.counter.FreshForce(), D, ss.UsageDecl, S),
			},
		})
		cs = cs.Merge(t.value(v))
	}

	return cs
}

func (t traverser) array(arr []any) ss.Constraints {
	cs := ss.Constraints{}

	S := ss.NewVariable(t.counter.FreshForce(), ss.BindingScope)
	if t.key != nil {
		cs = cs.Merge(ss.Constraints{
			AssociationKnown: []ss.AssociationKnown{ss.NewAssociationKnown(
				t.counter.FreshForce(), *t.key, S,
			)},
		})
	}
	for k, v := range arr {
		D := ss.NewIdentifier(fmt.Sprintf("%d", k), "?", 0, 0) //TODO: Add path info
		t.key = &D
		cs = cs.Merge(ss.Constraints{
			Usage: []ss.Usage{
				ss.NewUsage(t.counter.FreshForce(), D, ss.UsageDecl, S),
			},
		})
		cs = cs.Merge(t.value(v))
	}

	return cs
}

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		os.Exit(1)
	}
	request := ss.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	var counter ss.CounterService
	if request.CounterURL == nil {
		counter = ss.NewCounterServiceMock()
	} else {
		counter = ss.NewCounterServiceImpl(*request.CounterURL)
	}
	var ctx ss.TypeContext
	if request.TypeContext != nil {
		j, err := os.ReadFile(*request.TypeContext)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ctx, err = ss.UnmarshalContext(j)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var root any
	err := json.Unmarshal([]byte(request.Code), &root)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	traverser := traverser{
		key:     nil,
		counter: counter,
		ctx:     ctx,
	}
	constraints := traverser.value(root)
	j, err := json.MarshalIndent(constraints, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(j))
}

func main() {
	Run()
}
