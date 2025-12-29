package models

import (
	"time"
)

type PhoneVerification struct {
	ID			uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	Phone		string		`json:"phone"` 
	Otp			string		`json:"otp"`
	ExpireAt	time.Time	`json:"expire_at"`
	CreatedAt 	time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
}