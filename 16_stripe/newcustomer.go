package main

import "fmt"

// Customer xx
type Customer struct {
	ID string
}

func getNewCustomerByPointer() *Customer {
	return &Customer{ID: "cus_123"}
}

func getNewCustomerByValue() Customer {
	return Customer{ID: "cus_234"}
}

func main() {

	customer1 := getNewCustomerByPointer()
	customer2 := getNewCustomerByValue()

	fmt.Printf("%T\n", customer1)
	fmt.Println(customer1)

	fmt.Printf("%T\n", customer2)
	fmt.Println(customer2)
}
