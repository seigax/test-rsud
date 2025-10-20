package request

import (
	"fmt"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type GetErrorMessageQuery struct {
	BasePaginateRequest
	model.ErrorMessage
}

func (query *GetErrorMessageQuery) GetOrderQuery() string {
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

type CreateErrorMessageRequest struct {
	Code            string `json:"code" validate:"required"`
	Type            string `json:"type" validate:"required"`
	ApplicationName string `json:"application_name" validate:"required"`
	Message         string `json:"message" validate:"required"`
	IsActiveFlag    string `json:"is_active_flag" validate:"enum=Y/N"`
	CreatedBy       uint   `json:"-"`
}

type UpdateErrorMessageRequest struct {
	ID              uint   `json:"-" validate:"required"`
	Code            string `json:"code" validate:"required"`
	Type            string `json:"type" validate:"required"`
	ApplicationName string `json:"application_name" validate:"required"`
	Message         string `json:"message" validate:"required"`
	IsActiveFlag    string `json:"is_active_flag" validate:"enum=Y/N"`
	UpdatedBy       uint   `json:"-"`
}
