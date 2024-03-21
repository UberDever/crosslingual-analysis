package main_test

import (
	"testing"
	translate "translate-go"
	ss "translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 5/server.go"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}
