package application

import (
	"github.com/Lenstack/farm_management/internal/core/services"
	desc "github.com/Lenstack/farm_management/pkg"
)

type MicroserviceServer struct {
	desc.UnimplementedMicroserviceServer
	MiddlewareApplication MiddlewareApplication
	AuthenticationService services.AuthenticationService
	RoleService           services.RoleService
	PermissionService     services.PermissionService
}

func NewMicroserviceServer(
	MiddlewareApplication MiddlewareApplication,
	AuthenticationService services.AuthenticationService,
	RoleService services.RoleService,
	PermissionService services.PermissionService,
) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication: MiddlewareApplication,
		AuthenticationService: AuthenticationService,
		RoleService:           RoleService,
		PermissionService:     PermissionService,
	}
}
