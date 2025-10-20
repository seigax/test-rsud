package request

import (
	"fmt"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type GetProvinceQuery struct {
	BasePaginateRequest
	model.Province
}

func (query *GetProvinceQuery) GetOrderQuery() string {
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

type CreateProvinceRequest struct {
	Name      string `json:"name" validate:"required"`
	CreatedBy uint   `json:"-"`
}

type UpdateProvinceRequest struct {
	ID        uint   `json:"-" validate:"required"`
	Name      string `json:"name" validate:"required"`
	UpdatedBy uint   `json:"-"`
}
