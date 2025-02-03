package services
import (
	"fmt"
	"github.com/HaseemKhattak01/stripe-integration/handlers"
	"github.com/HaseemKhattak01/stripe-integration/models"
)
type StripeService struct {
	StripeHandler *handlers.StripeHandler
}
func NewStripeService(handler *handlers.StripeHandler) *StripeService {
	return &StripeService{
		StripeHandler: handler,
	}
}
func (ss *StripeService) CreateCustomer(description string) (*models.Customer, error) {
	cus, err := ss.StripeHandler.CreateCustomer(description)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %v", err)
	}
	return &models.Customer{
		ID:          cus.ID,
		Description: cus.Description, 
		Email:       cus.Email,
	}, nil
}
func (ss *StripeService) CreateCharge(amount int64, currency string, customerID string) (*models.Charge, error) {
	ch, err := ss.StripeHandler.CreateCharge(amount, currency, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to create charge: %v", err)
	}
	return &models.Charge{
		ID:       ch.ID,
		Amount:   ch.Amount,
		Currency: string(ch.Currency),
		Customer: ch.Customer.ID,
	}, nil
}