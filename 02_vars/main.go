package main

import (
	"fmt"
)

func main() {
	// String
	// var a string = "David"
	// var a = "David"
	name := "David"

	// Int
	// var age int32 = 25
	age := 25

	// bool
	// var isCool bool = true
	isCool := false

	fmt.Println(name, age, isCool)

	fmt.Printf("%T\n", isCool)
}
