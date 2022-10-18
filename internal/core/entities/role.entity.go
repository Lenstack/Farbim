package entities

import "time"

const RoleTableName = "_roles"

type Role struct {
	Id            string
	Name          string
	PermissionsId string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
