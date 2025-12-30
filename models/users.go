package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             	uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	Phone 			string		`json:"phone"`
	IsPhoneVerified	*bool		`gorm:"default:false" json:"is_phone_verified"`
	CreatedAt		time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	Roles			[]Role 		`gorm:"many2many:user_roles"`
}


func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &PhoneVerification{}, &Alert{}, &DailyCheckIn{}, &Role{})
	return err
}

func AlterUsersTable(db *gorm.DB) error {
	if db.Migrator().HasColumn(&User{}, "is_otp_verified") {
		if err := db.Migrator().DropColumn(&User{}, "is_otp_verified"); err != nil {
			return err
		}
	}
	if !db.Migrator().HasColumn(&User{}, "is_phone_verified") {
		if err := db.Migrator().AddColumn(&User{}, "is_phone_verified"); err != nil {
			return err
		}
	}
	return nil
}