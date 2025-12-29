package models

import (
	"time"
)


type Alert struct {
	ID				uint		`gorm:"primaryKey;autoIncrement" json:"id"`
	ElderID			int 		`json:"elder_id"` 
	CaregiverID		int			`json:"caregiver_id"`
	AlertType		string		`json:"alert_type"`
	SentVia			string		`json:"sent_via"`
	CreatedAt		time.Time	`gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	Elder 			User 		`gorm:"foreignKey:ElderID;references:ID"`
	Caregiver 		User 		`gorm:"foreignKey:CaregiverID;references:ID"`
}
