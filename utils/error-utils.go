package utils
import (
	"fmt"
	"github.com/stripe/stripe-go/v76"
)
func HandleStripeError(err error) error {
	if stripeErr, ok := err.(*stripe.Error); ok {
		switch stripeErr.Code {
		case stripe.ErrorCodeCardDeclined:
			return fmt.Errorf("card was declined: %v", stripeErr.Error())
		case stripe.ErrorCodeExpiredCard:
			return fmt.Errorf("card is expired: %v", stripeErr.Error())
		default:
			return fmt.Errorf("stripe error: %v", stripeErr.Error())
		}
	}
	return err
}