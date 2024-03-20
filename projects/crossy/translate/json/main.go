package main

import (
	"encoding/json"
	"fmt"
	"os"
	"translate/shared"
)

// basically, this is a sum type of named tuples
type state interface{ stateVariant() }

type scope struct{}

func (scope) stateVariant() {}

type scopeRefDecl struct{}

func (scopeRefDecl) stateVariant() {}

type scopeType struct{}

func (scopeType) stateVariant() {}

func traverse(v any, s state, f func(v any, s state) bool) {
	f(v, s)
	switch val := v.(type) {
	case []any:
	case map[string]any:
		for _, v := range val {
			traverse(v, s, f)
		}
	}
}

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		os.Exit(1)
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	var counter shared.CounterService
	if request.CounterURL == nil {
		counter = &shared.CounterServiceMock{}
	} else {
		counter = shared.NewCounterServiceImpl(*request.CounterURL)
	}

	var root any
	err := json.Unmarshal([]byte(request.Code), &root)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	traverse(root, nil, func(v any, s state) bool {
		//NOTE: https://pkg.go.dev/encoding/json
		switch v.(type) {
		case float64:
			break
		case bool:
			break
		case string:
			break
		case nil:
			break
		case []any:
			break
		case map[string]any:
			break
		default:
			panic("Unreachable")
		}

		return true
	})

	_ = counter

	fmt.Println(string(""))
}

func main() {
	Run()
}
