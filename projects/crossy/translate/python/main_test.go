package main_test

import (
	"fmt"
	"testing"
	translate "translate-python"
	ss "translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 2/compute.py"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	err := ss.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := ss.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		fmt.Print(out)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
