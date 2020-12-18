package main

import (
	"fmt"
	"strconv"
)

// Person object
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
	address   *Address
}

// Address 1
type Address struct {
	street string
}

// greet with value receiver
func (p Person) greet() string {
	return "Hello, i am " + p.firstName + " " + p.lastName + "living at " + p.city + ", i'm " + strconv.Itoa(p.age)
}

// change value with pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func pointerReceiver(p *Person) {
	fmt.Println(p, p.firstName)
	fmt.Printf("pointer Receiver type: %T\n", p)
	fmt.Printf("pointer Receiver address: %p\n", p)
}

func main() {
	address := &Address{
		street: "ChangPing Load",
	}
	person1 := Person{
		firstName: "David",
		lastName:  "Zeng",
		city:      "Shanghai",
		gender:    "male",
		age:       25,
		address:   address,
	}

	fmt.Println(person1)
	fmt.Println(person1.lastName)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2 := &Person{
		firstName: "Tom",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "male",
		age:       30,
	}

	fmt.Println("person2 use attribute directly: " + person2.firstName)
	fmt.Printf("person2's type: %T\n", person2)
	fmt.Printf("person2 memory address is: %p\n", person2)
	pointerReceiver(person2)
}
