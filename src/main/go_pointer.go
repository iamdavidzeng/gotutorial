package main

import "fmt"

func main() {
	x := 7
	y := 7

	sumBySelf(&x)
	sumAgain(y)

	fmt.Println(x)
	fmt.Println(y)
}


func sumBySelf(x *int) {
	*x++
}


func sumAgain(x int) {
	x++
}