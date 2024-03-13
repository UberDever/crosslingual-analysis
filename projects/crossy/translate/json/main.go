package main

import (
	"encoding/json"
	"log"
	"os"
	"translate/shared"
)

func traverse(v any, f func(v any) bool) {
	f(v)
	switch val := v.(type) {
	case []any:
		for i := range val {
			traverse(val[i], f)
		}
	case map[string]any:
		for _, v := range val {
			traverse(v, f)
		}
	}
}

func Run() {
	if len(os.Args) < 2 {
		log.Print("No argument were provided to translator")
		os.Exit(1)
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	var root any
	err := json.Unmarshal([]byte(request.Code), &root)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	traverse(root, func(v any) bool {
		return true
	})

	// TODO: Make this id global to all translators somehow
	var id uint = 0
	c := []shared.Constraint{
		shared.NewUsage(
			shared.NewIdentifier(id+0, "a", "/some/path", 0, 1),
			shared.UsageDecl,
			shared.NewVariable(id+1, shared.BindingScope),
		),
	}
	j, err := json.Marshal(c)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Print(string(j))
}

func main() {
	Run()
}
