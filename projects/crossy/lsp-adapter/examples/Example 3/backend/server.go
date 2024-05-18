package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type counter struct {
	Count int `json:"counter"`
}

func (c *counter) increment() {
	c.Count++
}

func main() {
	counter := counter{}

	http.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.Method {
		case "POST":
			counter.increment()
		case "GET":
			data, err := json.Marshal(counter)
			if err != nil {
				return
			}
			w.Write(data)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	log.Println("Server started at :3333")
	http.ListenAndServe(":3333", nil)
}
