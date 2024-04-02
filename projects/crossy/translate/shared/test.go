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
	counter *uint
}

func NewCounterServiceMock() CounterServiceMock {
	return CounterServiceMock{counter: new(uint)}
}

func (c CounterServiceMock) Fresh() (uint, error) {
	tmp := *c.counter
	*c.counter++
	return tmp, nil
}

func (c CounterServiceMock) FreshForce() uint {
	i, err := c.Fresh()
	if err != nil {
		panic(err)
	}
	return i
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

func CompareJsonOutput(expected, actual string) error {
	opts := jsondiff.DefaultConsoleOptions()
	difference, out := jsondiff.Compare([]byte(expected), []byte(actual), &opts)
	if difference != jsondiff.FullMatch {
		return fmt.Errorf("%s", out)
	}
	return nil
}

func ExtractConstraintsFromFile(file string, translator func()) (string, error) {
	code, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(path.Join(dir, file))
	if err != nil {
		return "", err
	}
	args := map[string]any{
		"id":                0,
		"code":              string(code),
		"path":              abs,
		"type_context_path": ANCHOR_PATH + "evaluation/ontology.json",
	}
	j, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	request := TryParseArguments(string(j))

	argsJson, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	got := RunAsCommand([]string{"", string(argsJson)}, translator)
	return got, nil
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
