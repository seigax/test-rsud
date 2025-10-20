package request

import (
	"fmt"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type GetMenuQuery struct {
	BasePaginateRequest
	model.Menu
}

func (query *GetMenuQuery) GetOrderQuery() string {
	fieldMap := map[string]string{
		"created_at":   "created_at",
		"order_number": "order_number",
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

type CreateMenuRequest struct {
	ParentMenuID uint   `json:"parent_menu_id"`
	Name         string `json:"name" validate:"required"`
	Level        string `json:"level" validate:"required"`
	Url          string `json:"url" validate:"required"`
	Icon         string `json:"icon"`
	IsActiveFlag string `json:"is_active_flag" validate:"enum=Y/N"`
	CreatedBy    uint   `json:"-"`
}

type UpdateMenuRequest struct {
	ID           uint   `json:"-" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Level        string `json:"level" validate:"required"`
	Url          string `json:"url" validate:"required"`
	Icon         string `json:"icon"`
	IsActiveFlag string `json:"is_active_flag" validate:"enum=Y/N"`
	UpdatedBy    uint   `json:"-"`
}

type UpdateOrderMenuRequest struct {
	Menus []struct {
		MenuID      uint `json:"id" validate:"required"`
		OrderNumber uint `json:"order_number" validate:"required"`
	} `json:"menus" validate:"dive"`
	UpdatedBy uint `json:"-"`
}
