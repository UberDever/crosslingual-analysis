package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"
	makefile "translate-makefile"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 2/Makefile"

func TestSmoke(t *testing.T) {
	out := shared.TestAsCommand([]string{}, makefile.Run)
	if len(out) == 0 {
		t.Errorf("Expected non-empty error")
	}
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
	request := shared.NewArguments(0, string(code), &abs)
	json, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	out := shared.TestAsCommand([]string{"", string(json)}, makefile.Run)
	fmt.Print(out)
}
