package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 4)

	for i :=0; i < 4; i++ {
		slice1[i] = i+1
	}

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}
