package main

import (
	"fmt"
	"os"

	"github.com/iamdavidzeng/gotutorial/16_stripe/utils"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentmethod"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	// customerID := os.Getenv("STRIPE_CUSTOMER_ID")
	// if customerID == "" {
	// 	fmt.Println("Please set STRIPE_CUSTOMER_ID for this request.")
	// }
	// listCustomerCards(customerID)

	detachPaymentMethod()
}

func listCustomerCards(customerID string) {

	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerID),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}

	i := paymentmethod.List(params)

	for i.Next() {
		pm := i.PaymentMethod()

		fmt.Println(utils.MarshalIndent(pm))
	}
}

func detachPaymentMethod() {
	paymentMethodID := os.Getenv("STRIPE_PAYMENT_METHOD_ID")

	pm, _ := paymentmethod.Detach(
		paymentMethodID, nil)

	fmt.Println(utils.MarshalIndent(pm))
}
