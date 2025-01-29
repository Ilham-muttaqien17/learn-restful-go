package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)


type ValidationErrorResponse struct {
	Message string `json:"message"`;
	Errors map[string]string `json:"errors"`
}

func Validator[T interface{}](model T) (T, map[string][]string) {
	validate := validator.New()

	if err := validate.Struct(model); err != nil {
		errorMessages := ParseValidationError(err, model)

		return model, errorMessages 
	}

	return model, nil
}

func ParseValidationError(err error, model interface{}) map[string][]string {
	errors := make(map[string][]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := getJSONFieldName(model, fieldError.StructField())
			message := parseValidationErrorMessage(fieldError, fieldName)

			errors[fieldName] = append(errors[fieldName], message)
		}
	}

	return errors

}

func parseValidationErrorMessage(err validator.FieldError, field string) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("The %s value is required", field)
	case "email":
		return "Invalid email address"
	case "min":
		return fmt.Sprintf("The %s value is too low", field)
	case "max":
		return  fmt.Sprintf("The %s value is too long", field)
	default:
		return "Invalid value"
	}
}