package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Code         string         `json:"code"`
	ParentMenuID uint           `json:"parent_menu_id"`
	Name         string         `json:"name"`
	Level        string         `json:"level"`
	Url          string         `json:"url"`
	Icon         string         `json:"icon"`
	OrderNumber  uint           `json:"order_number"`
	IsActiveFlag string         `json:"is_active_flag"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    uint           `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"-"`

	Menu []Menu `json:"menus" gorm:"foreignKey:ID;references:ParentMenuID"`
}

func (Menu) TableName() string {
	return "menu"
}

func (menu *Menu) IsChild() bool {
	return menu.Level == "Child"
}
