package model

import (
	"time"

	"gorm.io/gorm"
)

type District struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CityID       uint           `json:"city_id"`
	Name         string         `json:"name"`
	Code         string         `json:"code"`
	IsActiveFlag string         `json:"is_active_flag"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    uint           `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"-"`

	City City `json:"city" gorm:"foreignKey:ID;references:CityID"`
}

func (District) TableName() string {
	return "district"
}
