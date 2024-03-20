package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

const DEFAULT_PORT = "63982"
const DEFAULT_INITIAL = 0

type count32 int32

func (c *count32) inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *count32) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func main() {
	port := DEFAULT_PORT
	var counter count32 = DEFAULT_INITIAL

	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch r.Method {
		case "GET":
			w.Write([]byte(fmt.Sprintf("{ \"value\": %d }", counter.get())))
			counter.inc()
		}
	})

	log.Println("Server started at :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
