package main

import "fmt"


func main() {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf("- g is type of %T and has value %v\n", g, g)
	}
}
