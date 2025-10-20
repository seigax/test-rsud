package model

import "time"

type User struct {
	ID                     uint       `json:"id" gorm:"primaryKey"`
	Code                   string     `json:"code"`
	Name                   string     `json:"name"`
	Email                  string     `json:"email"`
	PhotoURL               string     `json:"photo_url"`
	EncryptedPassword      string     `json:"-"`
	ChangePasswordAt       time.Time  `json:"change_password_at"`
	TncAcceptedAt          time.Time  `json:"tnc_accepted_at"`
	LoginWithBiometricFlag string     `json:"login_with_biometric_flag"`
	IsActiveFlag           string     `json:"is_active_flag"`
	CreatedAt              time.Time  `json:"created_at"`
	CreatedBy              uint       `json:"created_by"`
	UpdatedAt              time.Time  `json:"updated_at"`
	UpdatedBy              uint       `json:"updated_by"`
	DeletedAt              *time.Time `json:"-"`

	UserPhone []UserPhone `json:"user_phones" gorm:"foreignKey:UserID;references:ID"`
	UserRole  []UserRole  `json:"user_roles" gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "user"
}
