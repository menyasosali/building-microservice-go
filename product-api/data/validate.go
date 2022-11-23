package data

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func Validate(i interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(i)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
