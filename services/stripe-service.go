package services

import (
	"fmt"

	"github.com/HaseemKhattak01/stripe-integration/models"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/token"
)

type StripeService struct {
	client *client.API
}

func NewStripeService(stripeKey string) *StripeService {
	stripeClient := &client.API{}
	stripeClient.Init(stripeKey, nil)
	return &StripeService{client: stripeClient}
}

func (ss *StripeService) CreateCustomer(description, email string) (*models.Customer, error) {
	params := &stripe.CustomerParams{
		Description: stripe.String(description),
		Email:       stripe.String(email),
	}
	cus, err := ss.client.Customers.New(params)
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

func (ss *StripeService) GenerateToken(cardDetails models.CardDetails) (string, error) {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(cardDetails.CardNumber),
			ExpMonth: stripe.String(cardDetails.ExpMonth),
			ExpYear:  stripe.String(cardDetails.ExpYear),
			CVC:      stripe.String(cardDetails.CVC),
		},
	}
	tok, err := token.New(params)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tok.ID, nil
}