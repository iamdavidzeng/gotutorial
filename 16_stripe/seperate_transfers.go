package main

import (
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/transfer"
)

func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	// transferGroupID := "DAVIDZENG-GROUP-1"
	sourceTransaction := "ch_1H3CrPDZL4nYRnYWQKaESTJw"

	// params := &stripe.PaymentIntentParams{
	// 	Amount:             stripe.Int64(10000),
	// 	Currency:           stripe.String(string(stripe.CurrencyGBP)),
	// 	PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	// 	TransferGroup:      stripe.String(transferGroupID),
	// }
	// pi, _ := paymentintent.New(params)
	// fmt.Println(pi.ClientSecret)

	// Create first Transfer
	transferParams := &stripe.TransferParams{
		Amount:      stripe.Int64(700),
		Currency:    stripe.String(string(stripe.CurrencyGBP)),
		Destination: stripe.String(os.Getenv("FIRST_CONNECTED_ACCOUNT_ID")),
		// TransferGroup: stripe.String(transferGroupID),
		SourceTransaction: stripe.String(sourceTransaction),
	}
	transfer.New(transferParams)

	// Create second Transfer
	secondTransferParams := &stripe.TransferParams{
		Amount:      stripe.Int64(200),
		Currency:    stripe.String(string(stripe.CurrencyGBP)),
		Destination: stripe.String(os.Getenv("SECOND_CONNECTED_ACCOUNT_ID")),
		// TransferGroup: stripe.String(transferGroupID),
		SourceTransaction: stripe.String(sourceTransaction),
	}
	transfer.New(secondTransferParams)
}
