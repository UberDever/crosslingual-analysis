package shared

import (
	"encoding/json"
	"net/http"
)

type CounterService interface {
	Get() (int, error)
}

type CounterServiceImpl struct {
	url string
}

func NewCounterServiceImpl(url string) CounterServiceImpl {
	return CounterServiceImpl{
		url: url,
	}
}
func (c CounterServiceImpl) Get() (int, error) {
	r, err := http.Get(c.url)
	if err != nil {
		return -1, err
	}
	defer r.Body.Close()

	var counter any
	err = json.NewDecoder(r.Body).Decode(counter)
	if err != nil {
		return -1, err
	}

	obj, ok := counter.(map[string]any)
	if !ok {
		return -1, nil
	}

	val, ok := obj["value"].(float64)
	if !ok {
		return -1, nil
	}

	return int(val), nil
}
