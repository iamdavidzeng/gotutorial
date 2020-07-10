package main

import (
	"fmt"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main() {

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CustomerParams{}
	c, _ := customer.New(params)

	fmt.Println(c)
}
