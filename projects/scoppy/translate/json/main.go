package main

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("No argument were provided to translator")
		return
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	var root any
	err := json.Unmarshal([]byte(request.Code), &root)
	if err != nil {
		log.Fatal(err)
	}

	traverse(root, func(v any) bool {
		fmt.Println(v)
		return true
	})
}
