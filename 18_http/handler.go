package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var result map[string]interface{}

	json.NewDecoder(r.Body).Decode(&result)

	a, _ := json.Marshal(result)
	fmt.Println(string(a))

	response, _ := json.Marshal("hello, world!")

	fmt.Fprintf(w, string(response))
}

func main() {
	handler := MyHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: &handler,
	}
	fmt.Println("Server is listening on http://localhost:8081")
	server.ListenAndServe()
}
