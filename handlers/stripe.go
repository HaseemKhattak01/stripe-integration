package handlers
import (
	"fmt"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/client"
)
type StripeHandler struct {
	Client *client.API
}
func NewStripeHandler(apiKey string) (*StripeHandler, error) {
	stripeClient := &client.API{}
	stripeClient.Init(apiKey, nil)
	return &StripeHandler{
		Client: stripeClient,
	}, nil
}
func (sh *StripeHandler) CreateCustomer(description string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Description: stripe.String(description),
	}
	cus, err := sh.Client.Customers.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create customer: %v", err)
	}
	return cus, nil
}
func (sh *StripeHandler) CreateCharge(amount int64, currency string, customerID string) (*stripe.Charge, error) {
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Customer: stripe.String(customerID),
	}
	ch, err := sh.Client.Charges.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create charge: %v", err)
	}
	return ch, nil
}