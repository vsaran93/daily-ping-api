package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID             	uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	Phone 			string		`json:"phone"`
	IsPhoneVerified	*bool		`gorm:"default:false" json:"is_phone_verified"`
	CreatedAt		time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{}, &PhoneVerifications{})
	return err
}

func AlterUsersTable(db *gorm.DB) error {
	if db.Migrator().HasColumn(&Users{}, "is_otp_verified") {
		if err := db.Migrator().DropColumn(&Users{}, "is_otp_verified"); err != nil {
			return err
		}
	}
	if !db.Migrator().HasColumn(&Users{}, "is_phone_verified") {
		if err := db.Migrator().AddColumn(&Users{}, "is_phone_verified"); err != nil {
			return err
		}
	}
	return nil
}