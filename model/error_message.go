package model

import (
	"time"

	"gorm.io/gorm"
)

type ErrorMessage struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Code            string         `json:"code"`
	Type            string         `json:"type"`
	ApplicationName string         `json:"application_name"`
	Message         string         `json:"message"`
	IsActiveFlag    string         `json:"is_active_flag"`
	CreatedAt       time.Time      `json:"created_at"`
	CreatedBy       uint           `json:"created_by"`
	UpdatedAt       time.Time      `json:"updated_at"`
	UpdatedBy       uint           `json:"updated_by"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}

func (ErrorMessage) TableName() string {
	return "error_message"
}
