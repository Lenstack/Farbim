package entities

import "time"

const SessionTableName = "_sessions"

type Session struct {
	Id           string
	Type         string
	Token        string
	Blocked      bool
	ClientOrigin string
	ExpiredAt    string
	UserId       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
