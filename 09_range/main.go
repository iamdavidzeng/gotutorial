package main

import "fmt"

func main() {
	// array
	ids := []int{11, 22, 33, 44}
	for i, v := range ids {
		fmt.Println(i, v)
	}

	// slice
	newIds := ids[:len(ids)-1]
	for i, v := range newIds {
		fmt.Println(i, v)
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
