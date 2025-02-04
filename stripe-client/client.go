package stripeclient

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/client"
)

// StripeClient is a type that wraps the Stripe API client.
type StripeClient struct {
	API *client.API
}

// Client is the global Stripe client instance.
var Client *StripeClient

// InitClient initializes the Stripe client with the given API key.
func InitClient(apiKey string) {
	stripe.Key = apiKey
	api := &client.API{}
	api.Init(apiKey, nil)
	Client = &StripeClient{API: api}
}
