package entities

import "time"

type User struct {
	Id                     string
	Email                  string
	Password               string
	Token                  string
	LastResetSentAt        time.Time
	LastVerificationSentAt time.Time
	Verified               bool
	Created                time.Time
	Updated                time.Time
}
