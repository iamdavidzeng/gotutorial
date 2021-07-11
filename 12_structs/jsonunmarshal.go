package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Private struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	person := `{"first_name": "david", "last_name": "zeng"}`

	var p1 Private
	err := json.NewDecoder(bytes.NewBuffer([]byte(person))).Decode(&p1)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("p1: ", p1)

	var p2 map[string]interface{}
	err = json.NewDecoder(bytes.NewBuffer([]byte(person))).Decode(&p2)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("p2: ", p2)

	var p3 Private
	err = json.Unmarshal([]byte(person), &p3)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("p3: ", p3)

}
