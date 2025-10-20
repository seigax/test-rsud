package request

import (
	"fmt"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type GetRoleQuery struct {
	BasePaginateRequest
	model.Role
}

func (query *GetRoleQuery) GetOrderQuery() string {
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

type CreateRoleRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	Type         string `json:"type" validate:"required"`
	Platform     string `json:"platform" validate:"required"`
	IsActiveFlag string `json:"is_active_flag" validate:"enum=Y/N"`
	CreatedBy    uint   `json:"-"`
}

type UpdateRoleRequest struct {
	ID           uint   `json:"-" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	Type         string `json:"type" validate:"required"`
	Platform     string `json:"platform" validate:"required"`
	IsActiveFlag string `json:"is_active_flag" validate:"enum=Y/N"`
	UpdatedBy    uint   `json:"-"`
}
