package utils

import (
	"fmt"
	"github.com/stripe/stripe-go/v76"
)

func HandleStripeError(err error) error {
	if stripeErr, ok := err.(*stripe.Error); ok {
		errMsg := getErrorMessage(stripeErr)
		return fmt.Errorf("%s: %v", errMsg, stripeErr.Error())
	}
	return err
}

func getErrorMessage(stripeErr *stripe.Error) string {
	switch stripeErr.Code {
	case stripe.ErrorCodeCardDeclined:
		return "card was declined"
	case stripe.ErrorCodeExpiredCard:
		return "card is expired"
	default:
		return "stripe error"
	}
}