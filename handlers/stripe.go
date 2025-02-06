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
	customerParams := &stripe.CustomerParams{Description: stripe.String(description)}
	cus, err := sh.createStripeCustomer(customerParams)
	if err != nil {
		return nil, err
	}

	paymentMethodParams := &stripe.PaymentMethodParams{
		Type: stripe.String("card"),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String("tok_visa"),
		},
	}
	pm, err := sh.createPaymentMethod(paymentMethodParams)
	if err != nil {
		return nil, err
	}

	err = sh.attachPaymentMethodToCustomer(pm.ID, cus.ID)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Created customer: %s\n", cus.ID)
	return cus, nil
}

func (sh *StripeHandler) createStripeCustomer(params *stripe.CustomerParams) (*stripe.Customer, error) {
	cus, err := customer.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}
	return cus, nil
}

func (sh *StripeHandler) createPaymentMethod(params *stripe.PaymentMethodParams) (*stripe.PaymentMethod, error) {
	pm, err := paymentmethod.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}
	return pm, nil
}

func (sh *StripeHandler) attachPaymentMethodToCustomer(paymentMethodID, customerID string) error {
	_, err := paymentmethod.Attach(paymentMethodID, &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customerID),
	})
	if err != nil {
		return fmt.Errorf("failed to attach payment method to customer: %w", err)
	}
	return nil
}
