package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	accountID := os.Args[1]
	a, _ := account.GetByID(accountID, nil)

	response, _ := json.Marshal(a)
	fmt.Println(string(response))
}
