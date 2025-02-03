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
	customerName := fmt.Sprintf("Test Customer %d", time.Now().UnixNano())
	customer, err := stripeClient.CreateCustomer(customerName)
	if err != nil {
		log.Fatalf("Failed to create customer: %v", err)
	}
	fmt.Printf("Created customer: %v\n", customer.ID)
}