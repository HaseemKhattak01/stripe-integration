package services

import (
	"fmt"

	"github.com/HaseemKhattak01/stripe-integration/handlers"
	"github.com/HaseemKhattak01/stripe-integration/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type StripeService struct {
	Handler *handlers.StripeHandler
}

func NewStripeService(handler *handlers.StripeHandler) *StripeService {
	return &StripeService{Handler: handler}
}

func (ss *StripeService) CreateCustomer(description string) (*models.Customer, error) {
	cus, err := ss.Handler.CreateCustomer(description)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %w", err)
	}
	return &models.Customer{
		ID:          cus.ID,
		Description: cus.Description,
		Email:       cus.Email,
	}, nil
}
func (ss *StripeService) CreatePaymentIntent(amount int64, currency, customerID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Customer: stripe.String(customerID),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}
	return pi, nil
}
