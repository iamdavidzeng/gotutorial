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
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	handler := MyHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	fmt.Println("Server is listening on http://localhost:8080")
	server.ListenAndServe()
}
