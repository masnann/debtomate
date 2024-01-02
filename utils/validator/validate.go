package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

func ValidateLoginRequest(req *LoginRequest) error {
	// Custom validation: password tidak boleh mengandung spasi
	validate.RegisterValidation("noSpace", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return !regexp.MustCompile(`\s`).MatchString(password)
	})

	err := validate.Struct(req)
	if err != nil {
		var customErrors []string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' is required", err.Field()))
			case "min":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' must be at least %s characters long", err.Field(), err.Param()))
			case "email":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' must be a valid email address", err.Field()))
			case "noSpace":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' cannot contain spaces", err.Field()))
			default:
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' validation failed with tag '%s'", err.Field(), err.Tag()))
			}
		}

		return fmt.Errorf("Validation error: %s", customErrors)
	}
	return nil
}