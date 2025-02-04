package models

type Customer struct {
	ID          string
	Description string
	Email       string
}

type CardDetails struct {
	CardNumber string `json:"card_number" validate:"required"`
	ExpMonth   string `json:"exp_month" validate:"required"`
	ExpYear    string `json:"exp_year" validate:"required"`
	CVC        string `json:"cvc" validate:"required"`
}
