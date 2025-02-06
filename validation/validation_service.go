package validation

import (
	"errors"
	"regexp"

	"github.com/HaseemKhattak01/stripe-integration/models"
)

type ValidationService struct{}

func NewValidationService() ValidationService {
	return ValidationService{}
}

func (vs *ValidationService) ValidateCardDetails(cardDetails models.CardDetails) error {
	if err := vs.validateField(cardDetails.CardNumber, vs.isValidCardNumber, "card number"); err != nil {
		return err
	}
	if err := vs.validateField(cardDetails.ExpMonth, vs.isValidMonth, "expiry month"); err != nil {
		return err
	}
	if err := vs.validateField(cardDetails.ExpYear, vs.isValidYear, "expiry year"); err != nil {
		return err
	}
	if err := vs.validateField(cardDetails.CVC, vs.isValidCVC, "CVC"); err != nil {
		return err
	}
	return nil
}

func (vs *ValidationService) ValidateCardNumber(cardNumber string) error {
	if err := vs.validateField(cardNumber, vs.isValidCardNumber, "card number"); err != nil {
		return err
	}
	return nil
}

func (vs *ValidationService) validateField(value string, validator func(string) bool, fieldName string) error {
	if value == "" {
		return errors.New(fieldName + " is required")
	}
	if !validator(value) {
		return errors.New("invalid " + fieldName)
	}
	return nil
}

func (vs *ValidationService) isValidCardNumber(number string) bool {
	re := regexp.MustCompile(`^\d{16}$`)
	return re.MatchString(number)
}

func (vs *ValidationService) isValidMonth(month string) bool {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])$`)
	return re.MatchString(month)
}

func (vs *ValidationService) isValidYear(year string) bool {
	re := regexp.MustCompile(`^([0-9]{4}|[0-9]{2})$`)
	return re.MatchString(year)
}

func (vs *ValidationService) isValidCVC(cvc string) bool {
	re := regexp.MustCompile(`^\d{3,4}$`)
	return re.MatchString(cvc)
}
