package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) CreateRole(ctx context.Context, request *pkg.CreateRoleRequest) (*pkg.CreateRoleResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	roleId, err := ms.RoleService.CreateRole(request.Name, request.PermissionsId)
	if err != nil {
		return nil, err
	}
	return &pkg.CreateRoleResponse{Message: "create role has been successfully", RoleId: roleId}, nil
}
