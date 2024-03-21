package shared

import (
	"encoding/json"
	"net/http"
)

type CounterService interface {
	Fresh() (uint, error)
	FreshForce() uint
}

type CounterServiceImpl struct {
	url string
}

func NewCounterServiceImpl(url string) CounterServiceImpl {
	return CounterServiceImpl{
		url: url,
	}
}
func (c CounterServiceImpl) Fresh() (uint, error) {
	r, err := http.Get(c.url)
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	var counter any
	err = json.NewDecoder(r.Body).Decode(counter)
	if err != nil {
		return 0, err
	}

	obj, ok := counter.(map[string]any)
	if !ok {
		return 0, nil
	}

	val, ok := obj["value"].(float64)
	if !ok {
		return 0, nil
	}

	return uint(val), nil
}

func (c CounterServiceImpl) FreshForce() uint {
	i, err := c.Fresh()
	if err != nil {
		panic(err)
	}
	return i
}
