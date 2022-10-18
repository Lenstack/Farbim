package entities

import (
	"time"
)

const UserTableName = "_users"

type User struct {
	Id              string
	Email           string
	Password        string
	TokenKey        string
	Verified        bool
	RolesId         string
	Code            string
	LastResetSentAt string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
