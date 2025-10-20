package model

import "time"

type UserVerification struct {
	ID                          uint       `json:"id" gorm:"primaryKey"`
	VerificationCode            string     `json:"verification_code"`
	UserID                      int        `json:"user_id"`
	VerificationTypeID          int        `json:"verification_type_id"`
	CommunicationDeviceTypeCode string     `json:"communication_device_type_code"`
	VerificationStatusFlag      bool       `json:"verification_status_flag"`
	ExpiredAt                   time.Time  `json:"expired_at"`
	CreatedAt                   time.Time  `json:"created_at"`
	CreatedBy                   int        `json:"created_by"`
	UpdatedAt                   time.Time  `json:"updated_at"`
	UpdatedBy                   int        `json:"updated_by"`
	DeletedAt                   *time.Time `json:"deleted_at"`
}

func (UserVerification) TableName() string {
	return "user_verification"
}
