package entities

import "time"

const FarmTableName = "farms"

type Farm struct {
	Id         string
	Name       string
	CategoryId string
	UserId     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
