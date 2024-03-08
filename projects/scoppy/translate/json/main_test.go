package main_test

import (
	"testing"
	json_translate "translate-json"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 5/weather.json"

func TestSmoke(t *testing.T) {
	out := shared.TestAsCommand([]string{}, json_translate.Run)
	if len(out) == 0 {
		t.Errorf("Expected non-empty error")
	}
}
