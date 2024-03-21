package main_test

import (
	"testing"
	translate "translate-makefile"
	ss "translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 2/Makefile"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	err := ss.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := ss.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		_ = out
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
