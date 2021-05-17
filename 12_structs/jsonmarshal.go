package main

import (
	"encoding/json"
	"fmt"
)

type NewPerson struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	david := &NewPerson{
		Name:   "david",
		Gender: "male",
	}
	out1, _ := json.Marshal(david)
	fmt.Println(string(out1))

	out2, _ := json.MarshalIndent(david, "", "  ")
	fmt.Println(string(out2))
}
