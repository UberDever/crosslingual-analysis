package main_test

import (
	"fmt"
	"testing"
	makefile "translate-makefile"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 2/Makefile"

func TestSmoke(t *testing.T) {
	out := shared.RunAsCommand([]string{}, makefile.Run)
	if len(out) == 0 {
		t.Errorf("Expected non-empty error")
	}
}

func TestEvaluation(t *testing.T) {
	err := shared.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := shared.RunAsCommand([]string{"", string(argsJson)}, makefile.Run)
		fmt.Print(out)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
