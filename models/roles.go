package models

import (
	"time"
)


type Role struct {
	ID			uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	Name		string 		`json:"name"`
	CreatedAt	time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	Users		[]User		`gorm:"many2many:user_roles"`
}