package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName string
}


func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}


func main() {
	//1-struct as a value type
	var person1 Person
	person1.firstName = "Chris"
	person1.lastName = "Woodward"
	upPerson(&person1)
	fmt.Printf("The name of the person1 is %s %s\n", person1.firstName, person1.lastName)

//	2-struct as a pointer:

	person2 := new(Person)
	person2.firstName = "Chris"
	person2.lastName = "Woodward"
	(*person2).lastName = "Woodward"
	upPerson(person2)
	fmt.Printf("The name of the person2 is %s %s\n", person2.firstName, person2.lastName)

//	3-struct as a literal
	person3 := &Person{"Chris", "Woodward"}
	upPerson(person3)
	fmt.Printf("The name if the person3 is %s %s\n", person3.firstName, person3.lastName)
}
