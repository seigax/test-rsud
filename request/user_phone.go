package request

import (
	"fmt"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type GetUserPhoneQuery struct {
	BasePaginateRequest
	model.UserPhone
}

func (query *GetUserPhoneQuery) GetOrderQuery() string {
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

type ReqAddPhone struct {
	Phone  string `json:"phone" validate:"required"`
	UserId uint   `json:"-"`
}

type ReqUpdatePhone struct {
	Phone  string `json:"phone" validate:"required"`
	UserId uint   `json:"-"`
	Id     uint   `json:"-"`
}
