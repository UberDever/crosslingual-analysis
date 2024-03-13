package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"translate/shared"
)

func jsonAst(code string) string {
	cmd := []string{"python3", "json_ast.py", code}
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.ExitError:
			log.Print(string(e.Stderr))
		}
		os.Exit(69)
	}
	return string(out)
}

func Run() {
	if len(os.Args) < 2 {
		log.Print("No argument were provided to translator")
		os.Exit(69)
	}
	request := shared.TryParseArguments(os.Args[1])
	if request == nil {
		return
	}
	codeJson := jsonAst(request.Code)
	var root any
	err := json.Unmarshal([]byte(codeJson), &root)
	if err != nil {
		log.Print(err)
		os.Exit(69)
	}

	log.Print(root)
}

func main() {
	Run()
}
