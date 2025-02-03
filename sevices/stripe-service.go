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
