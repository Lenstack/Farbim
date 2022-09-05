package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	*gorm.Model
	Email                  string
	Password               string
	Token                  string
	LastResetSentAt        time.Time
	LastVerificationSentAt time.Time
	Verified               bool
}
