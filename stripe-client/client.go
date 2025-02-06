package stripeclient

import (
	"errors"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/client"
)

type StripeClient struct {
	API *client.API
}

var Client *StripeClient

func InitClient(apiKey string) error {
	if apiKey == "" {
		return errors.New("STRIPE_API_KEY is not set")
	}
	stripe.Key = apiKey
	api := client.New(apiKey, nil)
	Client = &StripeClient{API: api}
	return nil
}
