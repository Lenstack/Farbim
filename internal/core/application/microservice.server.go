package application

import (
	"github.com/Lenstack/farm_management/internal/core/services"
	desc "github.com/Lenstack/farm_management/pkg"
)

type MicroserviceServer struct {
	MiddlewareApplication     MiddlewareApplication
	UserApplication           UserApplication
	AuthenticationApplication AuthenticationApplication
	desc.UnimplementedMicroserviceServer
	//Grpc Services
	AuthenticationService services.AuthenticationService
	UserService           services.UserService
	ProfileService        services.ProfileService
	FarmService           services.FarmService
	CategoryService       services.CategoryService
}

func NewMicroserviceServer(
	middlewareApplication MiddlewareApplication,
	userApplication UserApplication,
	authenticationApplication AuthenticationApplication,
	//Grpc Services
	AuthenticationService services.AuthenticationService,
	UserService services.UserService,
	ProfileService services.ProfileService,
	FarmService services.FarmService,
	CategoryService services.CategoryService,

) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication:     middlewareApplication,
		UserApplication:           userApplication,
		AuthenticationApplication: authenticationApplication,
		//Grpc Services
		AuthenticationService: AuthenticationService,
		UserService:           UserService,
		ProfileService:        ProfileService,
		FarmService:           FarmService,
		CategoryService:       CategoryService,
	}
}
