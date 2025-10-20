package model

import (
	"time"

	"gorm.io/gorm"
)

type SystemParameter struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Code          string         `json:"code"`
	ParameterName string         `json:"parameter_name"`
	DataType      string         `json:"data_type"`
	Message       string         `json:"message"`
	IsActiveFlag  string         `json:"is_active_flag"`
	CreatedAt     time.Time      `json:"created_at"`
	CreatedBy     uint           `json:"created_by"`
	UpdatedAt     time.Time      `json:"updated_at"`
	UpdatedBy     uint           `json:"updated_by"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

func (SystemParameter) TableName() string {
	return "system_parameter"
}
