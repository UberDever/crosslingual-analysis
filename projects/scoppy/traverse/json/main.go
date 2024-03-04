package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the source code to compile")
		os.Exit(-1)
	}
	source := os.Args[1]

	var root any
	err := json.Unmarshal([]byte(source), &root)
	if err != nil {
		log.Fatal(err)
	}

	traverse(root, func(v any) bool {
		fmt.Println(v)
		return true
	})
}
