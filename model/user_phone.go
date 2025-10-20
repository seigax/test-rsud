package model

import (
	"time"

	"gorm.io/gorm"
)

type UserPhone struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	UserID       uint           `json:"user_id"`
	PhoneNumber  string         `json:"phone_number"`
	IsActiveFlag string         `json:"is_active_flag"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    uint           `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (UserPhone) TableName() string {
	return "user_phone"
}
