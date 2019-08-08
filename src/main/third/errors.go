package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")


func main() {
	fmt.Printf("errors: %v", errNotFound)
}
