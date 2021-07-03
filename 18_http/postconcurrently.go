package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func post(addr string) *http.Response {
	body, _ := json.Marshal(map[string]string{"name": "david"})

	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(fmt.Errorf("err: %s", err))
	}

	return resp
}

func main() {

	addrs := os.Args[1:]

	c := make(chan *http.Response, len(addrs))
	quit := make(chan struct{})

	for _, addr := range addrs {
		go func(addr string) {
			c <- post(addr)
		}(addr)
	}

	go func() {
		for i := 0; i < len(addrs); i++ {
			fmt.Println(<-c)
		}
		close(quit)
	}()

	<-quit
}
