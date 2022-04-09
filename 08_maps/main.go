package main

import "fmt"

func main() {

	// Use make declare map
	// emails := make(map[string]string)

	// emails["David"] = "david@gmail.com"
	// emails["Bob"] = "bob@gmail.com"

	// Declare map and assign value
	emails := map[string]string{"David": "david.gmail.com", "Bob": "bob@gmail.com"}
	fmt.Println(emails)

	fmt.Println(len(map[string]int{"a": 2}))
}
