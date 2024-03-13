package main_test

import (
	"testing"
	translate "translate-xml"
	"translate/shared"
)

const MAIN_PATH = "../../../../evaluation/Example 3/CSharp/CSharp.proj"

func TestSmoke(t *testing.T) {
	shared.RunAsCommand([]string{"_", ""}, translate.Run)
}
