package shared

import (
	"io"
	"os"
)

func TestAsCommand(args []string, run func()) string {
	oldArgs := os.Args
	os.Args = args
	defer func() { os.Args = oldArgs }()

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	run()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout
	return string(out)
}
