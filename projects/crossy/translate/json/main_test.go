package main_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"testing"
	translate "translate-json"
	ss "translate/shared"
)

const MAIN_PATH = ss.ANCHOR_PATH + "evaluation/Example 5/weather.json"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	code, err := os.ReadFile(MAIN_PATH)
	if err != nil {
		t.Fatal(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	abs, err := filepath.Abs(path.Join(dir, MAIN_PATH))
	if err != nil {
		t.Fatal(err)
	}
	args := map[string]any{
		"id":                0,
		"code":              string(code),
		"path":              abs,
		"type_context_path": ss.ANCHOR_PATH + "evaluation/type_context.json",
	}
	j, err := json.Marshal(args)
	if err != nil {
		t.Fatal(err)
	}
	request := ss.TryParseArguments(string(j))

	err = ss.RunOnFile(*request, func(argsJson []byte) error {
		out := ss.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		_ = out

		dot, err := ss.ToDot(out)
		if err != nil {
			switch e := err.(type) {
			case *exec.ExitError:
				t.Fatal(string(e.Stderr))
			default:
				t.Fatal(e)
			}
		}
		t.Fatal(dot)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
