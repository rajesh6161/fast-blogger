package validators

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
	XValidator struct {
		validator *validator.Validate
	}
)

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	var validate = validator.New()
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func Validator(data interface{}) *fiber.Map {
	myValidator := &XValidator{
		validator: validator.New(),
	}
	if errs := myValidator.Validate(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		log.Printf("Validation error: %v", errMsgs)
		return &fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": strings.Join(errMsgs, " and "),
		}
	}
	return nil
}
