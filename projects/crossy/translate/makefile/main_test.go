package main_test

import (
	"fmt"
	"testing"
	translate "translate-makefile"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 2/Makefile"

func TestSmoke(t *testing.T) {
	shared.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	err := shared.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := shared.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		fmt.Print(out)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
