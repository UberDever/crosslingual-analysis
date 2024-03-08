package main

import (
	"encoding/json"
	"fmt"
	"os"
	"translate/shared"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("No argument were provided to translator")
		return
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	result := shared.NewResult(request.Id, []shared.Constraint{}, []shared.Unrecognized{
		shared.NewUnrecognized(*request.Path, 0, uint(len(request.Code)), request.Code),
	})
	json, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
}

func main() {
	Run()
}
