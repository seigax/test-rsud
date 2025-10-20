package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	ErrDateStrID string = "{0} bukanlah format YYYY-MM-DD"
	ErrDateStrEN string = "{0} not in format YYYY-MM-DD"
)

func Date(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	_, err := time.Parse("2006-01-02", value)
	return err == nil

}
