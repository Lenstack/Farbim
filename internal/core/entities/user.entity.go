package entities

import (
	"time"
)

const UserTableName = "users"

type User struct {
	Id                     string    `json:"Id,omitempty"`
	Email                  string    `json:"Email,omitempty" validate:"email,required"`
	Password               string    `json:"Password,omitempty" validate:"required"`
	AccessToken            string    `json:"AccessToken,omitempty"`
	RefreshToken           string    `json:"RefreshToken,omitempty"`
	LastResetSentAt        string    `json:"LastResetSentAt,omitempty"`
	LastVerificationSentAt string    `json:"LastVerificationSentAt,omitempty"`
	Verified               bool      `json:"Verified,omitempty"`
	Roles                  string    `json:"Roles"`
	CreatedAt              time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt              time.Time `json:"UpdatedAt,omitempty"`
}
