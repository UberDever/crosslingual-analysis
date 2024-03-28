package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type JSON = map[string]interface{}

func read_weather(path string) (JSON, error) {
	byteValue, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result JSON
	json.Unmarshal(byteValue, &result)
	return result, nil
}

func main() {

	js, err := read_weather("weather.json")
	if err != nil {
		println("Failed to open weather")
		os.Exit(-1)
	}

	response := ""
	if js["main"].(map[string]interface{})["temp"].(float64) >= 300 {
		response = "hot"
	} else {
		response = "cold"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p>"))
		w.Write([]byte("Today outside is "))
		w.Write([]byte(response))
		w.Write([]byte("</p>"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
