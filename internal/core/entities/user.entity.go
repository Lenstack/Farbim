package entities

import (
	"time"
)

const UserTableName = "users"

type User struct {
	Id                     string    `json:"Id,omitempty"`
	Email                  string    `json:"Email,omitempty" validate:"email,required"`
	Password               string    `json:"Password,omitempty" validate:"required"`
	Token                  string    `json:"Token,omitempty"`
	LastResetSentAt        time.Time `json:"LastResetSentAt,omitempty"`
	LastVerificationSentAt time.Time `json:"LastVerificationSentAt,omitempty"`
	Verified               bool      `json:"Verified,omitempty"`
	CreatedAt              time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt              time.Time `json:"UpdatedAt,omitempty"`
}
