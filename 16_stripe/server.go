package main

import (
	"html/template"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/setupintent"
)

// WalletData use to load for html
type WalletData struct {
	ClientSecret string
}

// Customer represent existing user in Stripe Platform
type Customer struct {
	ID string
}

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	customer := Customer{ID: os.Getenv("CUSTOMER")}

	cardWalletTmpl := template.Must(template.ParseFiles("views/card_wallet.html"))

	http.HandleFunc("/card-wallet", func(w http.ResponseWriter, r *http.Request) {
		params := &stripe.SetupIntentParams{
			Customer: stripe.String(customer.ID),
		}
		intent, err := setupintent.New(params)
		if err != nil {
			panic(err)
		}
		data := WalletData{
			ClientSecret: intent.ClientSecret,
		}
		cardWalletTmpl.Execute(w, data)
	})

	http.ListenAndServe(":3000", nil)
}
