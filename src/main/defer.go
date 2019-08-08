package main

import "fmt"

func main() {
	func1()
}


func func1() {
	fmt.Printf("In function1 at the top\n")
	defer func2()
	fmt.Printf("In function1 at the bottom\n")
}


func func2() {
	fmt.Printf("Deferred until the end of the calling function!")
}
