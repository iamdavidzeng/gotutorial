package main

import (
	"os"
	"fmt"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/payout"
)


func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	params := &stripe.PayoutParams{
		Amount: stripe.Int64(100),
		Currency: stripe.String(string(stripe.CurrencyGBP)),
	}

	po, _ := payout.New(params)
	
	fmt.Println(po)
}