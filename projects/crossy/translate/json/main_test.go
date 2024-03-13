package main_test

import (
	"fmt"
	"testing"
	json_translate "translate-json"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 5/weather.json"

func TestSmoke(t *testing.T) {
	out := shared.RunAsCommand([]string{}, json_translate.Run)
	if len(out) == 0 {
		t.Errorf("Expected non-empty error")
	}
}

func TestEvaluation(t *testing.T) {
	err := shared.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := shared.RunAsCommand([]string{"", string(argsJson)}, json_translate.Run)
		fmt.Print(out)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
