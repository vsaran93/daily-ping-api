package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID             	uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	Phone 			string		`json:"phone"`
	IsOtpVerified	*bool		`gorm:"default:false" json:"is_otp_verified"`
	CreatedAt		time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}