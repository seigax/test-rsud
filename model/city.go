package model

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	ProvinceID   uint           `json:"province_id"`
	Name         string         `json:"name"`
	Code         string         `json:"code"`
	IsActiveFlag string         `json:"is_active_flag"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    uint           `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"-"`

	Province Province `json:"province" gorm:"foreignKey:ID;references:ProvinceID"`
}

func (City) TableName() string {
	return "city"
}
