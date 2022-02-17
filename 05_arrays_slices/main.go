package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr1 := [2]string{"Apple", "Orange"}
	fmt.Println(reflect.TypeOf(fruitArr1))
	fmt.Println(len(fruitArr1))
	fmt.Println(fruitArr1)

	// Slices
	fruitArr2 := []string{"Apple", "Orange", "Banana", "Cherry"}
	fmt.Println(reflect.TypeOf(fruitArr2))
	fmt.Println(len(fruitArr2))
	fmt.Println(fruitArr2[:4])

	fruitArr3 := [4]int{}
	fruitArr3[0] = 1
	fmt.Println(fruitArr3)
}
