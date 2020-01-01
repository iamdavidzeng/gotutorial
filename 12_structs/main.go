package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greet with value reciever
func (p Person) greet() string {
	return "Hello, i am " + p.firstName + " " + p.lastName + "living at " + p.city + ", i'm " + strconv.Itoa(p.age)
}

// change value with pointer reciever
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person1 := Person{
		firstName: "David",
		lastName:  "Zeng",
		city:      "Shanghai",
		gender:    "male",
		age:       25,
	}

	fmt.Println(person1)
	fmt.Println(person1.lastName)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
