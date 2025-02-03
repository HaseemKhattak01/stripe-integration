package services

import (
	"fmt"

	"github.com/HaseemKhattak01/stripe-integration/handlers"
	"github.com/HaseemKhattak01/stripe-integration/models"
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

func (ss *StripeService) CreateCharge(amount int64, currency, customerID string) (*models.Charge, error) {
	ch, err := ss.Handler.CreateCharge(amount, currency, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to create charge: %w", err)
	}
	return &models.Charge{
		ID:       ch.ID,
		Amount:   ch.Amount,
		Currency: string(ch.Currency),
		Customer: ch.Customer.ID,
	}, nil
}