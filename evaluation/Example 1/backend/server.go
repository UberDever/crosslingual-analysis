package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	counter := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		byteValue, err := os.ReadFile("../frontend/index.html")
		if err != nil {
			println("Cannot open index page")
			os.Exit(-1)
		}
		w.Write(byteValue)
	})
	http.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3333")
		switch r.Method {
		case "POST":
			counter += 1
		case "GET":
			w.Write([]byte(fmt.Sprintf(`{"count": %d}`, counter)))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	log.Println("Server started at :3333")
	http.ListenAndServe(":3333", nil)
}
