package main_test

import (
	"testing"
	translate "translate-go"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 5/server.go"

func TestSmoke(t *testing.T) {
	shared.RunAsCommand([]string{"_", ""}, translate.Run)
}
