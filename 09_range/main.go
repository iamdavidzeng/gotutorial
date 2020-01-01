package main

import "fmt"

func main() {
	// array
	ids := []int{11, 22, 33, 44}
	for _, i := range ids {
		fmt.Println(i)
	}

	// map
	emails := map[string]string{
		"David": "david@gmail.com",
		"Bob":   "bob@gmail.com",
	}

	for k, v := range emails {
		fmt.Println(k, v)
	}
}
