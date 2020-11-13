package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/transfer"
)

// Use TransferGroupID bind Charge and Transfers, Maybe failed because of insufficient funds
func separateTransferWithTransferGroup() {

	timestampString := strconv.FormatInt(time.Now().Unix(), 10)

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(10000),
		Currency:           stripe.String(string(stripe.CurrencyGBP)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		TransferGroup:      stripe.String(timestampString),
	}
	pi, _ := paymentintent.New(params)
	fmt.Println(pi.ClientSecret)

	// Create first Transfer
	transferParams := &stripe.TransferParams{
		Amount:        stripe.Int64(700),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Destination:   stripe.String(os.Getenv("FIRST_CONNECTED_ACCOUNT_ID")),
		TransferGroup: stripe.String(timestampString),
	}
	firstTransfer, _ := transfer.New(transferParams)
	fmt.Println(firstTransfer)

	// Create second Transfer
	secondTransferParams := &stripe.TransferParams{
		Amount:        stripe.Int64(200),
		Currency:      stripe.String(string(stripe.CurrencyGBP)),
		Destination:   stripe.String(os.Getenv("SECOND_CONNECTED_ACCOUNT_ID")),
		TransferGroup: stripe.String(timestampString),
	}
	secondTransfer, _ := transfer.New(secondTransferParams)
	fmt.Println(secondTransfer)
}

// Use SourceTransaction
func separateTransferWithSourceTransaction() {

	sourceTransaction := os.Getenv("SOURCE_TRANSACTION")

	firstTransferParams := &stripe.TransferParams{
		Amount:            stripe.Int64(700),
		Currency:          stripe.String(string(stripe.CurrencyGBP)),
		SourceTransaction: stripe.String(sourceTransaction),
		Destination:       stripe.String(os.Getenv("FIRST_CONNECTED_ACCOUNT_ID")),
	}
	firstTransfer, _ := transfer.New(firstTransferParams)
	fmt.Println(firstTransfer)

	secondTransferParams := &stripe.TransferParams{
		Amount:            stripe.Int64(200),
		Currency:          stripe.String(string(stripe.CurrencyGBP)),
		SourceTransaction: stripe.String(sourceTransaction),
		Destination:       stripe.String(os.Getenv("SECOND_CONNECTED_ACCOUNT_ID")),
	}
	secondTransfer, _ := transfer.New(secondTransferParams)
	fmt.Println(secondTransfer)
}

// Use metadat in Transfer
func applyMetadataInTransfer() {
	sourceTransaction := os.Getenv("SOURCE_TRANSACTION")

	transferParams := &stripe.TransferParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"payment_item_ids": "[1, 2, 3]",
			},
		},
		Amount:            stripe.Int64(100),
		Currency:          stripe.String(string(stripe.CurrencyGBP)),
		SourceTransaction: stripe.String(sourceTransaction),
		Destination:       stripe.String(os.Getenv("CONNECTED_ACCOUNT_ID")),
	}

	myTransfer, _ := transfer.New(transferParams)
	fmt.Println(myTransfer)
}

func main() {

	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	applyMetadataInTransfer()
}
