package stripeclient

import "github.com/stripe/stripe-go/v76"

func InitClient(apiKey string) {
	stripe.Key = apiKey
}