package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) UpdateUserRoles(ctx context.Context, request *pkg.UpdateUserRolesRequest) (*pkg.UpdateUserRolesResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	userId, err := ms.RoleService.UpdateUserRole(request.UserId, request.RolesId)
	if err != nil {
		return nil, err
	}
	return &pkg.UpdateUserRolesResponse{Message: "roles has been updated successfully", UserId: userId}, nil
}
