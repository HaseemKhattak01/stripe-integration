package payment

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	const maxBodyBytes = 65536
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), "your-webhook-secret")
	if err != nil {
		http.Error(w, "Invalid webhook signature", http.StatusBadRequest)
		return
	}

	handleEvent(event, w)
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
