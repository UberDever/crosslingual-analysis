package main_test

import (
	"os/exec"
	"testing"
	translate "translate-json"
	"translate/shared"
)

const MAIN_PATH = shared.ANCHOR_PATH + "evaluation/Example 5/weather.json"

func TestSmoke(t *testing.T) {
	shared.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	err := shared.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := shared.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		dot, err := shared.ToDot(out)
		if err != nil {
			switch e := err.(type) {
			case *exec.ExitError:
				t.Fatal(string(e.Stderr))
			default:
				t.Fatal(e)
			}
		}
		t.Fatal(dot)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
