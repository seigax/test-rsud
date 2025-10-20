package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	RoleID    uint           `json:"role_id"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy uint           `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	UpdatedBy uint           `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Role Role `json:"role" gorm:"foreignKey:ID;references:RoleID"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type UserRoleWithRoleDetail struct {
	UserRole
	RoleName     string `json:"role_name"`
	RoleType     string `json:"role_type"`
	RolePlatform string `json:"role_platform"`
}
