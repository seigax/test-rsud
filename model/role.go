package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Code         string         `json:"code"`
	Description  string         `json:"description"`
	Type         string         `json:"type"`
	Platform     string         `json:"platform"`
	IsActiveFlag string         `json:"is_active_flag"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    uint           `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"-"`

	RoleMenu []RoleMenu `json:"role_menus" gorm:"foreignKey:RoleID;references:ID"`
}

func (Role) TableName() string {
	return "role"
}
