package main

import (
	"fmt"
	"os"

	"github.com/iamdavidzeng/gotutorial/16_stripe/utils"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	accountID := os.Args[1]
	a, _ := account.GetByID(accountID, nil)

	fmt.Println(utils.MarshalIndent(a))
}
