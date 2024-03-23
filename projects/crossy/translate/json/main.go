package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	ss "translate/shared"
)

type traverser struct {
	key *ss.Identifier

	path    string
	ctx     ss.TypeContext
	counter ss.CounterService
	decoder *json.Decoder
}

func (t traverser) value(value *json.Token) ss.Constraints {
	cs := ss.Constraints{}
	if !t.decoder.More() {
		return cs
	}

	var token json.Token
	if value != nil {
		token = *value
	} else {
		tt, err := t.decoder.Token()
		if err == io.EOF {
			return cs
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		token = tt
	}
	switch v := token.(type) {
	case json.Delim:
		if v == '{' {
			cs = cs.Merge(t.object())
			_, _ = t.decoder.Token() // '}'
		} else if v == '[' {
			cs = cs.Merge(t.array())
			_, _ = t.decoder.Token() // ']'
		} else {
			panic("Unreachable")
		}
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

func (t traverser) object() ss.Constraints {
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
	for t.decoder.More() {
		start := uint(t.decoder.InputOffset())
		token, err := t.decoder.Token()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		end := uint(t.decoder.InputOffset())
		k := token.(string)
		D := ss.NewIdentifier(k, t.path, start, end)
		t.key = &D
		cs = cs.Merge(ss.Constraints{
			Usage: []ss.Usage{
				ss.NewUsage(t.counter.FreshForce(), D, ss.UsageDecl, S),
			},
		})
		cs = cs.Merge(t.value(nil))
	}

	return cs
}

func (t traverser) array() ss.Constraints {
	cs := ss.Constraints{}

	S := ss.NewVariable(t.counter.FreshForce(), ss.BindingScope)
	if t.key != nil {
		cs = cs.Merge(ss.Constraints{
			AssociationKnown: []ss.AssociationKnown{ss.NewAssociationKnown(
				t.counter.FreshForce(), *t.key, S,
			)},
		})
	}
	index := 0
	for t.decoder.More() {
		start := uint(t.decoder.InputOffset())
		token, err := t.decoder.Token()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		end := uint(t.decoder.InputOffset())
		D := ss.NewIdentifier(fmt.Sprintf("%d", index), t.path, start, end)
		index++
		t.key = &D
		cs = cs.Merge(ss.Constraints{
			Usage: []ss.Usage{
				ss.NewUsage(t.counter.FreshForce(), D, ss.UsageDecl, S),
			},
		})
		cs = cs.Merge(t.value(&token))
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

	dec := json.NewDecoder(strings.NewReader(request.Code))
	base_path := "?"
	if request.Path != nil {
		base_path = filepath.Base(*request.Path)
	}

	traverser := traverser{
		key:     nil,
		path:    base_path,
		counter: counter,
		ctx:     ctx,
		decoder: dec,
	}
	constraints := traverser.value(nil)
	j, err := json.Marshal(constraints)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(j))
}

func main() {
	Run()
}
