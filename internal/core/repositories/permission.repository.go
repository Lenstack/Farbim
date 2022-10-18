package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
)

type IPermissionRepository interface {
	CreatePermission(permission entities.Permission) (permissionId string, err error)
}

type PermissionRepository struct {
	Database squirrel.StatementBuilderType
}

func (pr *PermissionRepository) CreatePermission(permission entities.Permission) (permissionId string, err error) {
	bq := pr.Database.
		Insert(entities.PermissionTableName).
		Columns("Id", "Service").
		Values(permission.Id, permission.Service).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&permissionId)
	if err != nil {
		return "", err
	}
	return permissionId, nil
}
