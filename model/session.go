package model

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID                       uint           `json:"id" gorm:"primaryKey"`
	UserID                   uint           `json:"user_id"`
	Token                    string         `json:"token"`
	IsLoginWithBiometricFlag string         `json:"is_login_with_biometric_flag"`
	CreatedAt                time.Time      `json:"created_at"`
	CreatedBy                uint           `json:"created_by"`
	UpdatedAt                time.Time      `json:"updated_at"`
	UpdatedBy                uint           `json:"updated_by"`
	DeletedAt                gorm.DeletedAt `json:"-"`
	ExpiredAt                time.Time      `json:"expired_at"`
}

func (Session) TableName() string {
	return "session"
}

func (session *Session) IsExpired() bool {
	return time.Now().After(session.ExpiredAt)
}
