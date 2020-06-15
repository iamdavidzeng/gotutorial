package main

import (
	"fmt"
	"os"
)

func main() {

	path := os.Getenv("PATH")
	fmt.Println(path)

	os.Exit(1)
}
