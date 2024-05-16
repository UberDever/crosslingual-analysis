package shared

import (
	"encoding/json"
	"fmt"
)

// TODO: Use source instead of Code and Path
type Arguments struct {
	Id         uint    `json:"id"`
	Code       string  `json:"code"`
	Path       *string `json:"path"`
	CounterURL *string `json:"counter_url"`
	Ontology   *string `json:"ontology"`
}

func TryParseArguments(arg string) *Arguments {
	var args Arguments
	err := json.Unmarshal([]byte(arg), &args)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &args
}
