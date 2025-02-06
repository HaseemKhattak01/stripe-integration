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
	fields := map[string]func(string) bool{
		"card number": vs.isValidCardNumber,
		"expiry month": vs.isValidMonth,
		"expiry year": vs.isValidYear,
		"CVC": vs.isValidCVC,
	}

	fieldValues := map[string]string{
		"card number": cardDetails.CardNumber,
		"expiry month": cardDetails.ExpMonth,
		"expiry year": cardDetails.ExpYear,
		"CVC": cardDetails.CVC,
	}

	for fieldName, validator := range fields {
		if err := vs.validateField(fieldValues[fieldName], validator, fieldName); err != nil {
			return err
		}
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
	return regexp.MustCompile(`^\d{16}$`).MatchString(number)
}

func (vs *ValidationService) isValidMonth(month string) bool {
	return regexp.MustCompile(`^(0[1-9]|1[0-2])$`).MatchString(month)
}

func (vs *ValidationService) isValidYear(year string) bool {
	return regexp.MustCompile(`^([0-9]{4}|[0-9]{2})$`).MatchString(year)
}

func (vs *ValidationService) isValidCVC(cvc string) bool {
	return regexp.MustCompile(`^\d{3,4}$`).MatchString(cvc)
}
