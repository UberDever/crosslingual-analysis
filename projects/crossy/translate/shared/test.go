package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/nsf/jsondiff"
)

const ANCHOR_PATH = "../../../../" // dir "crosslingual-analysis"

type CounterServiceMock struct {
	counter int
}

func (c *CounterServiceMock) Get() (int, error) {
	tmp := c.counter
	c.counter++
	return tmp, nil
}

func RunAsCommand(args []string, run func()) string {
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

func RunOnFile(codePath string, onTranslate func(argsJson []byte) error) error {
	code, err := os.ReadFile(codePath)
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	abs, err := filepath.Abs(path.Join(dir, codePath))
	if err != nil {
		return err
	}
	request := arguments{
		Id:   0,
		Code: string(code),
		Path: &abs,
	}
	argsJson, err := json.Marshal(request)
	if err != nil {
		return err
	}
	err = onTranslate(argsJson)
	if err != nil {
		return err
	}
	return nil
}

func CompareJsonOutput(expected, actual string) error {
	opts := jsondiff.DefaultConsoleOptions()
	difference, out := jsondiff.Compare([]byte(expected), []byte(actual), &opts)
	if difference != jsondiff.FullMatch {
		return fmt.Errorf("%s", out)
	}
	return nil
}

const toDotPath = ANCHOR_PATH + "evaluation/to_dot.py"

func ToDot(json string) (string, error) {
	input, err := os.CreateTemp("", "")
	if err != nil {
		return "", err
	}
	defer os.Remove(input.Name())
	_, err = input.Write([]byte(json))
	if err != nil {
		return "", err
	}

	output, err := os.CreateTemp("", "")
	if err != nil {
		return "", err
	}
	defer os.Remove(output.Name())

	cmd := []string{"python3", toDotPath,
		"-p", input.Name(),
		"-o", output.Name(),
		"-t", "json"}
	_, err = exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		return "", err
	}
	
	output.Close()
	out, err := os.ReadFile(output.Name())
	if err != nil {
		return "", err
	}

	return string(out), nil
}
