package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`
	PreferredFish []string  `json:"prefeered_fish,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

func main() {
	u := &User{
		Name:      "Sam",
		Password:  "password",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
