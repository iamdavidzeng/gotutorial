package main

import (
	"fmt"
	"go_record/16_stripe/utils"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	accountID := os.Args[1]
	a, _ := account.GetByID(accountID, nil)

	fmt.Println(utils.MarshalIndent(a))
}
