package utilities

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiBindError struct {
	Field string
	Msg   string
}

func generateErrorMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}

func ParseError(err error) []ApiBindError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiBindError, len(ve))
		for i, fe := range ve {
			out[i] = ApiBindError{fe.Field(), generateErrorMessage(fe.Tag())}
		}
		return out
	}
	return []ApiBindError{}
}
