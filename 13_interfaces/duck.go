package main

import "fmt"

type Duck interface {
	Quack()
}

type Cat struct{}

// 使用结构体
// func (c Cat) Quack() {
// 	fmt.Println("meow")
// }

// 使用结构体指针
func (c *Cat) Quack() {
	fmt.Println("meow")
}

func main() {
	// var d1 Duck = Cat{}
	// d1.Quack()

	var d2 Duck = &Cat{}
	d2.Quack()
}
