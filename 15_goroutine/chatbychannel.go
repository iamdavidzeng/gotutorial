package main

import (
	"fmt"
	"time"
)

// Foo describe methods
type Foo interface {
	greet() (string, error)
}

// Bar declare a abstract object
type Bar struct {
	firstName, lastName string
	age                 int64
}

func (b Bar) greet() (string, error) {
	return fmt.Sprintf("Hello, %v %v, you are %v years old now!", b.firstName, b.lastName, b.age), nil
}

func create(c chan Bar) error {

	go func() {
		time.Sleep(time.Second * 1)
		bar := Bar{
			firstName: "Foo",
			lastName:  "Bar",
			age:       18,
		}
		c <- bar
	}()

	return nil
}

func main() {

	channel := make(chan Bar)

	err := create(channel)

	bar := <-channel
	word, err := bar.greet()

	if err != nil {
		fmt.Println("Opps, something is wrong!")
	} else {
		fmt.Println(word)
	}
}
