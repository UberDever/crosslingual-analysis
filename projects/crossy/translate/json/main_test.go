package main_test

import (
	"testing"
	translate "translate-json"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 5/weather.json"

func TestSmoke(t *testing.T) {
	shared.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	err := shared.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := shared.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		t.Fatal(out)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
