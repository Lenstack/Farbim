package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type IRoleService interface {
	CreateRole(name string, permissionsId string) (roleId string, err error)
}

type RoleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(database squirrel.StatementBuilderType) *RoleService {
	return &RoleService{roleRepository: repositories.RoleRepository{Database: database}}
}

func (rs *RoleService) CreateRole(name string, permissionsId string) (roleId string, err error) {
	return rs.roleRepository.CreateRole(entities.Role{Id: uuid.New().String(), Name: name, PermissionsId: permissionsId})
}
