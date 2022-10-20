package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type IPermissionRepository interface {
	GetPermissions() (permissions []entities.Permission, err error)
	CreatePermission(permission entities.Permission) (permissionId string, err error)
}

type PermissionRepository struct {
	Database squirrel.StatementBuilderType
}

func (pr *PermissionRepository) GetPermissions() (permissions []entities.Permission, err error) {
	permission := entities.Permission{}
	bq := pr.Database.
		Select("Service", "Roles").
		From(entities.PermissionTableName)

	rows, err := bq.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&permission.Service, pq.Array(&permission.Roles))
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (pr *PermissionRepository) CreatePermission(permission entities.Permission) (permissionId string, err error) {
	bq := pr.Database.
		Insert(entities.PermissionTableName).
		Columns("Id", "Service", "Roles").
		Values(permission.Id, permission.Service, pq.Array(permission.Roles)).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&permissionId)
	if err != nil {
		return "", err
	}
	return permissionId, nil
}
