package shared

import (
	"encoding/json"
	"fmt"
)

type arguments struct {
	Id   uint    `json:"id"`
	Code string  `json:"code"`
	Path *string `json:"path"`
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

func NewArguments(id uint, code string, path *string) arguments {
	return arguments{
		Id:   id,
		Code: code,
		Path: path,
	}
}
