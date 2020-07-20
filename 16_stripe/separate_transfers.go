package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/transfer"
)

func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	transferGroupID := "202007201925"

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(10000),
		Currency:           stripe.String(string(stripe.CurrencyGBP)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		TransferGroup:      stripe.String(transferGroupID),
	}
	pi, _ := paymentintent.New(params)
	fmt.Println(pi.ClientSecret)

	// Create first Transfer
	transferParams := &stripe.TransferParams{
		Amount:        stripe.Int64(700),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Destination:   stripe.String(os.Getenv("FIRST_CONNECTED_ACCOUNT_ID")),
		TransferGroup: stripe.String(transferGroupID),
	}
	firstTransfer, _ := transfer.New(transferParams)
	fmt.Println(firstTransfer)

	// Create second Transfer
	secondTransferParams := &stripe.TransferParams{
		Amount:        stripe.Int64(200),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Destination:   stripe.String(os.Getenv("SECOND_CONNECTED_ACCOUNT_ID")),
		TransferGroup: stripe.String(transferGroupID),
	}
	secondTransfer, _ := transfer.New(secondTransferParams)
	fmt.Println(secondTransfer)
}
