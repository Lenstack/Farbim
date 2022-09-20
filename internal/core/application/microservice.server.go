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
}

func NewMicroserviceServer(
	middlewareApplication MiddlewareApplication,
	userApplication UserApplication,
	authenticationApplication AuthenticationApplication,
	//Grpc Services
	AuthenticationService services.AuthenticationService,
	UserService services.UserService,
	ProfileService services.ProfileService,

) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication:     middlewareApplication,
		UserApplication:           userApplication,
		AuthenticationApplication: authenticationApplication,
		//Grpc Services
		AuthenticationService: AuthenticationService,
		UserService:           UserService,
		ProfileService:        ProfileService,
	}
}
