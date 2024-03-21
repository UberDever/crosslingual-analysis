package main

import (
	"encoding/json"
	"fmt"
	"os"
	"translate/shared"
)

type traverser struct {
	counter shared.CounterService
}

func (t traverser) value(root any) shared.Constraints {
	cs := shared.Constraints{}
	switch v := root.(type) {
	case map[string]any:
		cs = cs.Merge(t.object(v))
	case []any:
		cs = cs.Merge(t.array(v))
	case float64:
		break
	case bool:
		break
	case string:
		break
	case nil:
		break
	}
	return cs
}

func (t traverser) object(v map[string]any) shared.Constraints {
	cs := shared.Constraints{}
	return cs
}

func (t traverser) array(v []any) shared.Constraints {
	cs := shared.Constraints{}
	return cs
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

	traverser := traverser{
		counter: counter,
	}
	constraints := traverser.value(root)
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
