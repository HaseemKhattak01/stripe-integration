package app

import (
	"log"

	"github.com/HaseemKhattak01/stripe-integration/config"
	"github.com/HaseemKhattak01/stripe-integration/handlers"
	"github.com/HaseemKhattak01/stripe-integration/payment"
	stripeclient "github.com/HaseemKhattak01/stripe-integration/stripe-client"
)

func RunApp() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	stripeclient.InitClient(cfg.StripeKey)

	stripeHandler := handlers.NewStripeHandler()

	cus, err := stripeHandler.CreateCustomer("Test Customer")
	if err != nil {
		log.Fatalf("Error in customer creation: %v", err)
	}
	log.Printf("Created customer: %v", cus.ID)

	paymentService := payment.NewPaymentService()
	paymentIntent, err := paymentService.CreatePaymentIntent(1000, "usd", cus.ID) // Example amount and currency
	if err != nil {
		log.Fatalf("Error creating payment intent: %v", err)
	}
	log.Printf("Created payment intent: %v", paymentIntent.ID)
}
