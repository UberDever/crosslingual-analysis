package main_test

import (
	"testing"
	translate "translate-cpp"
	ss "translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 4/lib.cpp"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	expected := ""
	got, err := ss.ExtractConstraintsFromFile(MAIN_PATH, translate.Run)
	if err != nil {
		t.Fatal(err)
	}
	if err := ss.CompareJsonOutput(expected, got); err != nil {
		t.Fatal(err)
	}
}
