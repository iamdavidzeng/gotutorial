package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var manager = map[string]*websocket.Conn{}

type IMMessage struct {
	Message string `json:"message"`
	From    string `json:"from"`
	To      string `json:"to"`
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User")
	fmt.Println(user, "connect to websocket successfully...")

	conn, err := upgrader.Upgrade(w, r, nil)

	manager[user] = conn

	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range manager {
		if k != user {
			if err := v.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s is online now.", user))); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	for {
		messageType, jsonMsg, err := manager[user].ReadMessage()
		msg := IMMessage{}
		json.Unmarshal(jsonMsg, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonMsg))

		if msg.To != "" {
			err := manager[msg.To].WriteMessage(messageType, []byte(msg.Message))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {

	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})

	http.HandleFunc("/ws", wsEndpoint)

	// listen and server on localhost:8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
