package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	ErrEnumStrID string = "{0} hanya boleh {1}"
	ErrEnumStrEN string = "{0} can only {1}"
)

func Enum(fl validator.FieldLevel) bool {
	enumString := fl.Param()
	value := fl.Field().String()
	enumSlice := strings.Split(enumString, "/")
	for _, v := range enumSlice {
		if value == v {
			return true
		}
	}
	return false
}
