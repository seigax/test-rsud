package request

import (
	"fmt"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type GetSystemParameterQuery struct {
	BasePaginateRequest
	model.SystemParameter
}

func (query *GetSystemParameterQuery) GetOrderQuery() string {
	fieldMap := map[string]string{
		"created_at": "created_at",
	}

	result := []string{}
	for _, s := range query.Sort {
		if len(s) == 0 {
			continue
		}

		order, key := "ASC", s
		if s[len(s)-1:] == "-" {
			order, key = "DESC", s[:len(s)-1]
		}

		fieldName, ok := fieldMap[key]
		if !ok {
			continue
		}

		result = append(result, fmt.Sprintf("%s %s", fieldName, order))
	}

	return strings.Join(result, ",")
}

type CreateSystemParameterRequest struct {
	ParameterName string `json:"parameter_name" validate:"required"`
	DataType      string `json:"data_type" validate:"required"`
	Message       string `json:"message" validate:"required"`
	IsActiveFlag  string `json:"is_active_flag" validate:"enum=Y/N"`
	CreatedBy     uint   `json:"-"`
}

type UpdateSystemParameterRequest struct {
	ID            uint   `json:"-" validate:"required"`
	ParameterName string `json:"parameter_name" validate:"required"`
	DataType      string `json:"data_type" validate:"required"`
	Message       string `json:"message" validate:"required"`
	IsActiveFlag  string `json:"is_active_flag" validate:"enum=Y/N"`
	UpdatedBy     uint   `json:"-"`
}
