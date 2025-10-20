package model

import (
	"time"

	"gorm.io/gorm"
)

type Village struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	DistrictID   uint           `json:"district_id"`
	Name         string         `json:"name"`
	PostalCode   string         `json:"postal_code"`
	Code         string         `json:"code"`
	IsActiveFlag string         `json:"is_active_flag"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    uint           `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    uint           `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"-"`

	District District `json:"district" gorm:"foreignKey:ID;references:DistrictID"`
}

func (Village) TableName() string {
	return "village"
}
