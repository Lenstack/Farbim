package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type IPermissionService interface {
	GetPermissions() (permissions []entities.Permission, err error)
	CreatePermission(serviceName string, rolesName []string) (permissionId string, err error)
}

type PermissionService struct {
	permissionRepository repositories.PermissionRepository
}

func NewPermissionService(database squirrel.StatementBuilderType) *PermissionService {
	return &PermissionService{permissionRepository: repositories.PermissionRepository{Database: database}}
}

func (ps *PermissionService) GetPermissions() (permissions []entities.Permission, err error) {
	return ps.permissionRepository.GetPermissions()
}

func (ps *PermissionService) CreatePermission(serviceName string, rolesName []string) (permissionId string, err error) {
	return ps.permissionRepository.CreatePermission(entities.Permission{Id: uuid.New().String(), Service: serviceName, Roles: rolesName})
}
