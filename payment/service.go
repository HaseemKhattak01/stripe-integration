package payment

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
)

const maxBodyBytes = 65536

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	if payload, err := io.ReadAll(r.Body); err != nil {
		http.Error(w, "Failed to read request body", http.StatusServiceUnavailable)
	} else if event, err := constructStripeEvent(payload, r.Header.Get("Stripe-Signature")); err != nil {
		http.Error(w, "Invalid webhook signature", http.StatusBadRequest)
	} else {
		handleEvent(event, w)
	}
}

func constructStripeEvent(payload []byte, signature string) (stripe.Event, error) {
	return webhook.ConstructEvent(payload, signature, "your-webhook-secret")
}

func handleEvent(event stripe.Event, w http.ResponseWriter) {
	switch event.Type {
	case "payment_intent.succeeded":
		fmt.Println("PaymentIntent was successful!")
	default:
		fmt.Printf("Unhandled event type: %s\n", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}
