package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello," + req.URL.Path[1:])
}


func main () {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
