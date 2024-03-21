package main_test

import (
	"os/exec"
	"testing"
	translate "translate-json"
	ss "translate/shared"
)

const MAIN_PATH = ss.ANCHOR_PATH + "evaluation/Example 5/weather.json"

func TestSmoke(t *testing.T) {
	ss.RunAsCommand([]string{"_", ""}, translate.Run)
}

func TestEvaluation(t *testing.T) {
	err := ss.RunOnFile(MAIN_PATH, func(argsJson []byte) error {
		out := ss.RunAsCommand([]string{"", string(argsJson)}, translate.Run)
		dot, err := ss.ToDot(out)
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
