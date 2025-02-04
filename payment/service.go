package payment

import (
	"fmt"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (ps *PaymentService) CreatePaymentIntent(amount int64, currency, customerID string) (*stripe.PaymentIntent, error) {
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

func (ps *PaymentService) ConfirmPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	pi, err := paymentintent.Confirm(paymentIntentID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to confirm payment intent: %w", err)
	}
	return pi, nil
}
