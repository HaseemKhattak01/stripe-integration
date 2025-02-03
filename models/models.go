package models
type Customer struct {
	ID          string
	Description string
	Email       string
}
type Charge struct {
	ID       string
	Amount   int64
	Currency string
	Customer string
}