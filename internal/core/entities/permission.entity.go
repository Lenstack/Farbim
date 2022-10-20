package entities

import "time"

const PermissionTableName = "_permissions"

type Permission struct {
	Id        string
	Service   string
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
