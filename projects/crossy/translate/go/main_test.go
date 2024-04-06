package main_test

import (
	"testing"
	translate "translate-go"
	ss "translate/shared"
)

const MAIN_PATH = ss.ANCHOR_PATH + "evaluation/Example 1/backend/server.go"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	expected := ""
	got, err := ss.ExtractConstraintsFromFile(MAIN_PATH, translate.Run)
	t.Fatal(got, err)
	if err != nil {
		t.Fatal(err)
	}
	if err := ss.CompareJsonOutput(expected, got); err != nil {
		t.Fatal(err)
	}
}
