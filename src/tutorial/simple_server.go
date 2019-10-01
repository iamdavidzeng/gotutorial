package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})

	// listen and server on localhost:8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
