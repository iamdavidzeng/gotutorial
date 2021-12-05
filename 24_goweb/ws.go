package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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

var oldMessages = []IMMessage{}

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

	for _, msg := range oldMessages {
		if msg.To == user {
			manager[user].WriteMessage(
				websocket.TextMessage,
				[]byte(msg.Message),
			)
		}
	}

	for {
		messageType, jsonMsg, err := manager[user].ReadMessage()
		msg := IMMessage{}
		json.Unmarshal(jsonMsg, &msg)
		oldMessages = append(oldMessages, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user, "got message: ", string(jsonMsg))

		if msg.To != "" {
			err = manager[msg.To].WriteMessage(messageType, []byte(jsonMsg))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsEndpoint)

	// listen and server on localhost:8081
	fmt.Println("Server is listening on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
