package app

import (
	"fmt"
	"log"
	"time"

	"github.com/HaseemKhattak01/stripe-integration/config"
	"github.com/HaseemKhattak01/stripe-integration/handlers"
)

func RunApp() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	stripeClient, err := handlers.NewStripeHandler(cfg.StripeKey)
	if err != nil {
		log.Fatalf("Failed to initialize Stripe client: %v", err)
	}

	if err := createAndLogCustomer(stripeClient); err != nil {
		log.Fatalf("Error in customer creation: %v", err)
	}
}

func createAndLogCustomer(stripeClient *handlers.StripeHandler) error {
	customerName := generateCustomerName()
	customer, err := stripeClient.CreateCustomer(customerName)
	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	log.Printf("Created customer: %v", customer.ID)
	return nil
}

func generateCustomerName() string {
	return fmt.Sprintf("Test Customer %d", time.Now().UnixNano())
}
