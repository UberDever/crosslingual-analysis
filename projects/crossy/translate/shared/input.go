package shared

import (
	"encoding/json"
	"fmt"
)

type arguments struct {
	Id         uint    `json:"id"`
	Code       string  `json:"code"`
	Path       *string `json:"path"`
	CounterURL *string `json:"counter_url"`
	Ontology   *string `json:"type_context_path"`
}

func TryParseArguments(arg string) *arguments {
	var args arguments
	err := json.Unmarshal([]byte(arg), &args)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &args
}
