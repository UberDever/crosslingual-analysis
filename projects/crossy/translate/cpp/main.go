package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"translate/shared"
)

func jsonAst(path string) string {
	cmd := []string{"clang++", "-x", "c++",
		"-Xclang", "-ast-dump-all=json", "-fsyntax-only", path}
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
	if request.Path == nil {
		log.Print("Path must be provided to this translator")
		os.Exit(69)
	}
	codeJson := jsonAst(*request.Path)
	var root any
	err := json.Unmarshal([]byte(codeJson), &root)
	if err != nil {
		log.Print(codeJson)
		log.Print(err)
		os.Exit(69)
	}

	log.Print(root)
}

func main() {
	Run()
}
