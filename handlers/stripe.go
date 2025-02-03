package handlers

import (
	stripeclient "github.com/HaseemKhattak01/stripe-integration/stripe-client"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/charge"
	"github.com/stripe/stripe-go/v76/client"
	"github.com/stripe/stripe-go/v76/customer"
)

type StripeHandler struct {
	Client *client.API
}

func NewStripeHandler(apiKey string) (*StripeHandler, error) {
	stripeclient.InitClient(apiKey)
	stripeClient := &client.API{}
	return &StripeHandler{Client: stripeClient}, nil
}

func (sh *StripeHandler) CreateCustomer(description string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{Description: stripe.String(description)}
	return customer.New(params)
}

func (sh *StripeHandler) CreateCharge(amount int64, currency, customerID string) (*stripe.Charge, error) {
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Customer: stripe.String(customerID),
	}
	return charge.New(params)
}
