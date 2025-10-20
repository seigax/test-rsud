package model

import (
	"time"

	"gorm.io/gorm"
)

type RoleMenu struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	RoleID    uint           `json:"role_id"`
	MenuID    uint           `json:"menu_id"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy uint           `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy uint           `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Menu Menu `json:"menu" gorm:"foreignKey:ID;references:MenuID"`
}

func (RoleMenu) TableName() string {
	return "role_menu"
}

type RoleMenuWithRoleUrl struct {
	RoleMenu
	MenuUrl string `json:"menu_url"`
}
