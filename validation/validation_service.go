package validation

import (
	"errors"
	"regexp"

	"github.com/HaseemKhattak01/stripe-integration/models"
)

type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (vs *ValidationService) ValidateCardDetails(cardDetails models.CardDetails) error {
	if err := validateField(cardDetails.CardNumber, isValidCardNumber, "card number"); err != nil {
		return err
	}
	if err := validateField(cardDetails.ExpMonth, isValidMonth, "expiry month"); err != nil {
		return err
	}
	if err := validateField(cardDetails.ExpYear, isValidYear, "expiry year"); err != nil {
		return err
	}
	if err := validateField(cardDetails.CVC, isValidCVC, "CVC"); err != nil {
		return err
	}
	return nil
}

func validateField(value string, validator func(string) bool, fieldName string) error {
	if value == "" {
		return errors.New(fieldName + " is required")
	}
	if !validator(value) {
		return errors.New("invalid " + fieldName)
	}
	return nil
}

func isValidCardNumber(number string) bool {
	re := regexp.MustCompile(`^\d{16}$`)
	return re.MatchString(number)
}

func isValidMonth(month string) bool {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])$`)
	return re.MatchString(month)
}

func isValidYear(year string) bool {
	re := regexp.MustCompile(`^([0-9]{4}|[0-9]{2})$`)
	return re.MatchString(year)
}

func isValidCVC(cvc string) bool {
	re := regexp.MustCompile(`^\d{3,4}$`)
	return re.MatchString(cvc)
}
