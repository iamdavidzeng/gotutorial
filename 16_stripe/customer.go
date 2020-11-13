package main

import (
	"fmt"
	"go_record/16_stripe/utils"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	getCustomer()
}

func getCustomer() {
	customerID := os.Getenv("STRIPE_CUSTOMER_ID")

	c, _ := customer.Get(customerID, nil)

	fmt.Println(utils.MarshalIndent(c))
}

func createCustomer() {

	createCustomerParams := &stripe.CustomerParams{
		Email: stripe.String("david.zeng@student.com"),
		Name:  stripe.String("david.zeng@student.com"),
	}
	response, _ := customer.New(createCustomerParams)

	fmt.Println(utils.MarshalIndent(response))
}

func updateCustomer() {
	updateCustomerParams := &stripe.CustomerParams{
		Email: stripe.String("david.zeng@student.com"),
	}
	customerID := os.Getenv("STRIPE_CUSTOMER_ID")
	if customerID == "" {
		fmt.Println("Please set STRIPE_CUSTOMER_ID before request.")
		return
	}
	response, _ := customer.Update(
		customerID,
		updateCustomerParams,
	)

	fmt.Println(utils.MarshalIndent(response))
}
