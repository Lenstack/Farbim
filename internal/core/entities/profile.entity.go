package entities

import "time"

const ProfileTableName = "_profiles"

type Profile struct {
	Id        string
	Name      string
	Avatar    string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
