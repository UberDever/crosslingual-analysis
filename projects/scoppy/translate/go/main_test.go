package main_test

import (
	"testing"
	golang "translate-go"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 5/server.go"

func TestSmoke(t *testing.T) {
	out := shared.RunAsCommand([]string{}, golang.Run)
	if len(out) == 0 {
		t.Errorf("Expected non-empty error")
	}
}
