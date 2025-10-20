package request

import (
	"fmt"
	"strings"
)

type GetUserQuery struct {
	BasePaginateRequest
}

func (query *GetUserQuery) GetOrderQuery() string {
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

type CreateUserRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	Roles           []uint `json:"roles_id" validate:"required,min=1"`
	IsActiveFlag    string `json:"is_active_flag" validate:"enum=Y/N"`
	CreatedBy       uint   `json:"-"`
}

type UpdateUserRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"min=6|min=0"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	Roles           []uint `json:"roles_id" validate:"required,min=1"`
	IsActiveFlag    string `json:"is_active_flag" validate:"enum=Y/N"`
	UpdatedBy       uint   `json:"-"`
	UserID          uint   `json:"-"`
}
