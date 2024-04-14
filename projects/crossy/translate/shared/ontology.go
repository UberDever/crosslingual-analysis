package shared

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/mitchellh/mapstructure"
)

type Ontology struct {
	TypesPath     string      `json:"types_path"`
	TemplatesPath string      `json:"templates_path"`
	counter       CounterService
	engine        *goja.Runtime
}

func loadIntoEngine(e *goja.Runtime, path string) error {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		if filepath.Ext(info.Name()) == ".js" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			val, err := e.RunString(string(data))
			if err != nil {
				return err
			}
			if !val.SameAs(goja.Undefined()) {
				return fmt.Errorf("haven't expected return value %v from engine", val)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func NewOntology(counter CounterService, path string) (c Ontology, err error) {
	c.counter = counter
	c.engine = goja.New()
	c.engine.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return
	}

	err = loadIntoEngine(c.engine, filepath.Join(filepath.Dir(path), c.TemplatesPath))
	if err != nil {
		return
	}

	err = loadIntoEngine(c.engine, filepath.Join(filepath.Dir(path), c.TypesPath))
	if err != nil {
		return
	}
	return
}

func (c Ontology) EvalTemplate(name string, counter CounterService, args ...any) (cs Constraints, result []any, err error) {
	var f func([]any) ([]any, error)
	fn := c.engine.Get(name)
	if fn == nil {
		err = fmt.Errorf("cannot find %s", name)
		return
	}
	err = c.engine.ExportTo(fn, &f)
	if err != nil {
		return
	}
	result, err = f(append([]any{counter}, args...))
	if err != nil {
		return
	}
	if len(result) < 1 {
		err = fmt.Errorf("expected return value in %s", name)
		return
	}
	dynCs, ok := result[0].(map[string]any)
	if !ok {
		err = fmt.Errorf("expected constraints as first return value, not %v", dynCs)
		return
	}
	err = mapstructure.Decode(dynCs, &cs)
	if err != nil {
		return
	}
	result = result[1:]

	return
}


type Type any

func (c Ontology) ConcreteType(name string) (t Type, err error) {
	var concreteType func(string) (string, error)
	fn := c.engine.Get("ConcreteType")
	if fn == nil {
		err = fmt.Errorf("cannot find %s", "ConcreteType")
		return
	}
	err = c.engine.ExportTo(fn, &concreteType)
	if err != nil {
		return
	}

	t, err = concreteType(name)
	if err != nil {
		return
	}
	return
}

func (c Ontology) NewType(ctor string, types ...Type) (t Type, err error) {
	for _, ty := range types {
		switch v := ty.(type) {
		case string:
			_, err = c.ConcreteType(v)
			if err != nil {
				return
			}
		}
	}
	var newType func(string, []Type) (any, error)
	fn := c.engine.Get("NewType")
	if fn == nil {
		err = fmt.Errorf("cannot find %s", "NewType")
		return
	}
	err = c.engine.ExportTo(fn, &newType)
	if err != nil {
		return
	}
	return newType(ctor, types)
}

