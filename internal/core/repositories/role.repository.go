package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
)

type IRoleRepository interface {
	CreateRole(role entities.Role) (roleId string, err error)
}

type RoleRepository struct {
	Database squirrel.StatementBuilderType
}

func (rr *RoleRepository) CreateRole(role entities.Role) (roleId string, err error) {
	bq := rr.Database.
		Insert(entities.RoleTableName).
		Columns("Id", "Name", "PermissionsId").
		Values(role.Id, role.Name, role.PermissionsId).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&roleId)
	if err != nil {
		return "", err
	}
	return roleId, nil
}
