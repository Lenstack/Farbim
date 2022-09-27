package application

import (
	"github.com/Lenstack/farm_management/internal/core/services"
	desc "github.com/Lenstack/farm_management/pkg"
)

type MicroserviceServer struct {
	desc.UnimplementedMicroserviceServer
	MiddlewareApplication MiddlewareApplication
	AuthenticationService services.AuthenticationService
	UserService           services.UserService
	ProfileService        services.ProfileService
	FarmService           services.FarmService
	CategoryService       services.CategoryService
}

func NewMicroserviceServer(
	MiddlewareApplication MiddlewareApplication,
	AuthenticationService services.AuthenticationService,
	UserService services.UserService,
	ProfileService services.ProfileService,
	FarmService services.FarmService,
	CategoryService services.CategoryService,

) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication: MiddlewareApplication,
		AuthenticationService: AuthenticationService,
		UserService:           UserService,
		ProfileService:        ProfileService,
		FarmService:           FarmService,
		CategoryService:       CategoryService,
	}
}
