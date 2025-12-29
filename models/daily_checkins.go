package models

import (
	"time"
)


type DailyCheckIn struct {
	ID				uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	ElderId			int 		`json:"elder_id"` 
	Status			string		`json:"status"`
	CreatedAt		time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	Elder 			User 		`gorm:"foreignKey:ElderID;references:ID"`
}
