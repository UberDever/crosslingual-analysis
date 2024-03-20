package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/nsf/jsondiff"
)

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
	difference, out := jsondiff.Compare([]byte(expected), []byte(actual), nil)
	if difference != jsondiff.FullMatch {
		return fmt.Errorf("%s", out)
	}
	return nil
}
