package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// PAN symbol validator: regex to match the format (XXXXX1234X)
func PANSymbol(fl validator.FieldLevel) bool {
	pan := fl.Field().String()
	r := regexp.MustCompile(`^[A-Za-z]{5}[0-9]{4}[A-Za-z]{1}$`)
	return r.MatchString(pan)
}

// Mobile validator: ensures it's exactly 10 digits
func Mobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	r := regexp.MustCompile(`^\d{10}$`)
	return r.MatchString(mobile)
}

// Register custom validators to be used globally
func RegisterValidators(validate *validator.Validate) {
	validate.RegisterValidation("pansymbol", PANSymbol)
	validate.RegisterValidation("mobile", Mobile)
}
