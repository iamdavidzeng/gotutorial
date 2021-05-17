package main

import "fmt"

type Foo struct {
	bar string
}

// wrong way
var foo1 *Foo

// correct way
var foo2 *Foo = &Foo{}

func (f *Foo) baz() error {
	fmt.Println("i am baz...")
	return nil
}

func (f Foo) buz() error {
	fmt.Println("i am buz...")
	return nil
}

func main() {
	foo1.baz()
	foo2.baz()
	foo2.buz()
	foo1.buz()
}
