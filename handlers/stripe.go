package handlers

import (
	"fmt"

	stripeclient "github.com/HaseemKhattak01/stripe-integration/stripe-client"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
)

type StripeHandler struct {
	Client *stripeclient.StripeClient
}

func NewStripeHandler() *StripeHandler {
	return &StripeHandler{Client: stripeclient.Client}
}

func (sh *StripeHandler) CreateCustomer(description string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{Description: stripe.String(description)}
	cus, err := customer.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %v", err)
	}
	fmt.Printf("Created customer: %s\n", cus.ID) // Log the customer ID
	return cus, nil
}