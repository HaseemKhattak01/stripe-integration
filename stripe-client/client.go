package stripeclient

import (
	"errors"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/client"
)

type StripeClient struct {
	API *client.API
}

var Client *client.API

func InitClient(apiKey string) error {
	if apiKey == "" {
		return errors.New("STRIPE_API_KEY is not set")
	}
	stripe.Key = apiKey
    Client = client.New(apiKey, nil)
    return nil
}
