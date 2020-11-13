package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(1000),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Customer:      stripe.String(os.Getenv("STRIPE_CUSTOMER_ID")),
		PaymentMethod: stripe.String(os.Getenv("STRIPE_PAYMENT_METHOD")),
		Confirm:       stripe.Bool(true),
		OffSession:    stripe.Bool(true),
	}

	_, err := paymentintent.New(params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			fmt.Printf("Error code: %v", stripeErr.Code)

			paymentIntentID := stripeErr.PaymentIntent.ID
			paymentIntent, _ := paymentintent.Get(paymentIntentID, nil)

			fmt.Printf("PI: %v", paymentIntent.ID)
		}
	}
}
