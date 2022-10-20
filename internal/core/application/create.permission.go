package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) CreatePermission(ctx context.Context, request *pkg.CreatePermissionRequest) (*pkg.CreatePermissionResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	permissionId, err := ms.PermissionService.CreatePermission(request.ServiceName, request.RolesName)
	if err != nil {
		return nil, err
	}
	return &pkg.CreatePermissionResponse{Message: "create permission has been successfully", PermissionId: permissionId}, nil
}
