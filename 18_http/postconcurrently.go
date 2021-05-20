package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func post(addr string, c chan<- bool) error {
	body, _ := json.Marshal(map[string]string{"name": "david"})

	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	fmt.Println(resp)

	c <- true

	return nil
}

func main() {

	addrs := os.Args[1:]

	c := make(chan bool)

	for _, addr := range addrs {
		go func(addr string) {
			post(addr, c)
		}(addr)
	}

	<-c
}
