package main

import (
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/v71/transfer"
)

func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(10000),
		Currency:           stripe.String(string(stripe.CurrencyGBP)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	pi, _ := paymentintent.New(params)

	// Create first Transfer
	transferParams := &stripe.TransferParams{
		Amount:        stripe.Int64(7000),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Destination:   stripe.String(os.Getenv("FIRST_CONNECTED_ACCOUNT_ID")),
		TransferGroup: stripe.String(*&pi.TransferGroup),
	}
	tr1, _ := transfer.New(transferParams)

	// Create second Transfer
	secondTransferParams := &stripe.TransferParams{
		Amount:        stripe.Int64(2000),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Destination:   stripe.String(os.Getenv("SECOND_CONNECTED_ACCOUNT_ID")),
		TransferGroup: stripe.String(*&pi.TransferGroup),
	}
	tr2, _ := transfer.New(secondTransferParams)
}
