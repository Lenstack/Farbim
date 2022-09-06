package entities

import (
	"time"
)

type User struct {
	Id                     string    `json:"Id" gorm:"primaryKey"`
	Email                  string    `json:"Email"`
	Password               string    `json:"Password"`
	Token                  string    `json:"Token"`
	LastResetSentAt        time.Time `json:"LastResetSentAt"`
	LastVerificationSentAt time.Time `json:"LastVerificationSentAt"`
	Verified               bool      `json:"Verified"`
}
