package request

import (
	"fmt"
	"strings"
)

type GetVillageQuery struct {
	BasePaginateRequest
	DistrictID uint
}

func (query *GetVillageQuery) GetOrderQuery() string {
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

type CreateVillageRequest struct {
	Name       string `json:"name" validate:"required"`
	DistrictID uint   `json:"district_id" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	CreatedBy  uint   `json:"-"`
}

type UpdateVillageRequest struct {
	ID         uint   `json:"-" validate:"required"`
	Name       string `json:"name" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	UpdatedBy  uint   `json:"-"`
}
