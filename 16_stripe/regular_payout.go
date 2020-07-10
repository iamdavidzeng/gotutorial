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
		// Specify Destination to Payout, Otherwise use default external account instead
		// Destination: stripe.String(os.Getenv("BANK_ACCOUNT")),
	}
	// Payout for Connected Account
	// params.SetStripeAccount(os.Getenv("FIRST_CONNECTED_ACCOUNT_ID"))

	po, _ := payout.New(params)
	
	fmt.Println(po)
}