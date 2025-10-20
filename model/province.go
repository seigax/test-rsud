package model

import (
	"time"

	"gorm.io/gorm"
)

type Province struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name"`
	Code          string         `json:"code"`
	TotalCity     uint           `json:"total_city"`
	TotalDistrict uint           `json:"total_district"`
	TotalVillage  uint           `json:"total_village"`
	IsActiveFlag  string         `json:"is_active_flag"`
	CreatedAt     time.Time      `json:"created_at"`
	CreatedBy     uint           `json:"created_by"`
	UpdatedAt     time.Time      `json:"updated_at"`
	UpdatedBy     uint           `json:"updated_by"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

func (Province) TableName() string {
	return "province"
}
