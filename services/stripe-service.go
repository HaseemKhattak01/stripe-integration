package services

import (
	"fmt"

	"github.com/HaseemKhattak01/stripe-integration/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

type StripeService struct {
	client *client.API
}

func NewStripeService(stripeKey string) StripeService {
	return StripeService{client: client.New(stripeKey, nil)}
}

func (ss *StripeService) CreateCustomer(description, email string) (*models.Customer, error) {
	cus, err := ss.client.Customers.New(&stripe.CustomerParams{
		Description: stripe.String(description),
		Email:       stripe.String(email),
	})
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
	pi, err := ss.client.PaymentIntents.New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Customer: stripe.String(customerID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}
	return pi, nil
}

func (ss *StripeService) GenerateToken(cardDetails models.CardDetails) (string, error) {
	tok, err := ss.client.Tokens.New(&stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(cardDetails.CardNumber),
			ExpMonth: stripe.String(cardDetails.ExpMonth),
			ExpYear:  stripe.String(cardDetails.ExpYear),
			CVC:      stripe.String(cardDetails.CVC),
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tok.ID, nil
}
