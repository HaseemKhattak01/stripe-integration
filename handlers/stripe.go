package handlers

import (
	"fmt"

	stripeclient "github.com/HaseemKhattak01/stripe-integration/stripe-client"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/paymentmethod"
)

type StripeHandler struct {
	Client *stripeclient.StripeClient
}

func NewStripeHandler() *StripeHandler {
	return &StripeHandler{Client: stripeclient.Client}
}

func (sh *StripeHandler) CreateCustomer(description string) (*stripe.Customer, error) {
	cus, err := customer.New(&stripe.CustomerParams{Description: stripe.String(description)})
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}

	pm, err := paymentmethod.New(&stripe.PaymentMethodParams{
		Type: stripe.String("card"),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String("tok_visa"),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}

	if _, err := paymentmethod.Attach(pm.ID, &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(cus.ID),
	}); err != nil {
		return nil, fmt.Errorf("failed to attach payment method to customer: %w", err)
	}

	fmt.Printf("Created customer: %s\n", cus.ID)
	return cus, nil
}
