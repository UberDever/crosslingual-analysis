package main_test

import (
	"testing"
	xml_translate "translate-xml"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 3/CSharp/CSharp.proj"

func TestSmoke(t *testing.T) {
	out := shared.TestAsCommand([]string{}, xml_translate.Run)
	if len(out) == 0 {
		t.Errorf("Expected non-empty error")
	}
}
