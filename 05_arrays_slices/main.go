package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// Slices
	fruitArr := []string{"Apple", "Orange", "Banana", "Cherry"}

	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:2])

}
