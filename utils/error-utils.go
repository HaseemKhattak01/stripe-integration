package utils

import (
	"fmt"
	"github.com/stripe/stripe-go/v76"
)

func HandleStripeError(err error) error {
	if stripeErr, ok := err.(*stripe.Error); ok {
		var errMsg string
		switch stripeErr.Code {
		case stripe.ErrorCodeCardDeclined:
			errMsg = "card was declined"
		case stripe.ErrorCodeExpiredCard:
			errMsg = "card is expired"
		default:
			errMsg = "stripe error"
		}
		return fmt.Errorf("%s: %v", errMsg, stripeErr.Error())
	}
	return err
}