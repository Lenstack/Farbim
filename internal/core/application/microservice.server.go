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
	AuthenticationService services.AuthenticationService
	UserService           services.UserService
}

func NewMicroserviceServer(
	middlewareApplication MiddlewareApplication,
	userApplication UserApplication,
	authenticationApplication AuthenticationApplication,
	AuthenticationService services.AuthenticationService,
	UserService services.UserService,

) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication:     middlewareApplication,
		UserApplication:           userApplication,
		AuthenticationApplication: authenticationApplication,
		AuthenticationService:     AuthenticationService,
		UserService:               UserService,
	}
}
