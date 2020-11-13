package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/setupintent"
)

// WalletData use to load for html
type WalletData struct {
	StripePublicKey string
	ClientSecret    string
}

// Customer represent existing user in Stripe Platform
type Customer struct {
	ID string
}

func main() {
	stripe.Key = os.Getenv("STRIPE_ACCOUNT_SECRET_KEY")

	customerID := os.Getenv("STRIPE_CUSTOMER_ID")
	if customerID == "" {
		fmt.Println("Please set STRIPE_CUSTOMER_ID for this request.")
	}

	customer := Customer{ID: customerID}

	setupIntentServer(customer)

	paymentIntentServer(customer)

	fmt.Println("Server listen on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func setupIntentServer(customer Customer) {
	cardWalletTmpl := template.Must(template.ParseFiles("views/setup_intent_card.html"))

	http.HandleFunc("/setup-intent", func(w http.ResponseWriter, r *http.Request) {
		params := &stripe.SetupIntentParams{
			Customer: stripe.String(customer.ID),
			Params: stripe.Params{
				Metadata: map[string]string{
					"name": "iamdavidzeng",
				},
			},
		}
		intent, err := setupintent.New(params)
		if err != nil {
			panic(err)
		}
		data := WalletData{
			StripePublicKey: os.Getenv("STRIPE_ACCOUNT_PUBLIC_KEY"),
			ClientSecret:    intent.ClientSecret,
		}
		cardWalletTmpl.Execute(w, data)
	})
}

func paymentIntentServer(customer Customer) {
	cardWalletTmpl := template.Must(template.ParseFiles("views/payment_intent_card.html"))

	http.HandleFunc("/payment-intent", func(w http.ResponseWriter, r *http.Request) {
		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(1000),
			Currency: stripe.String(string(stripe.CurrencyGBP)),
			Customer: stripe.String(customer.ID),
			Params: stripe.Params{
				Metadata: map[string]string{
					"name": "iamdavidzeng",
				},
			},
		}
		intent, err := paymentintent.New(params)
		if err != nil {
			panic(err)
		}
		data := WalletData{
			StripePublicKey: os.Getenv("STRIPE_ACCOUNT_PUBLIC_KEY"),
			ClientSecret:    intent.ClientSecret,
		}
		cardWalletTmpl.Execute(w, data)
	})
}
