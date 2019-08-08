package main

import "fmt"

type innerS struct {
	int1 int
	int2 int
}

type outerS struct {
	b int
	c float32
	int // anonymous field
	innerS //anonymous field
}


func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.int1 = 5
	outer.int2 = 10


	outer2 := outerS{6, 7.5, 60, innerS{5,10}}
	fmt.Println("outer2 is:", outer2)
}
