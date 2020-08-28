package main

import (
	"fmt"
)

func main() {
	const name, age = "David", 26
	s := fmt.Sprintf("%s is %d years old.", name, age)

	fmt.Println(s)
}
