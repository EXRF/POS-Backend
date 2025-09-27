package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorMessage(fe validator.FieldError) string {
	field := fe.Field()
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
	case "e164":
		return fmt.Sprintf("%s must be a valid international phone number", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
