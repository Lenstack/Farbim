package application

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignUp(_ context.Context, request *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {
	user := entities.User{Email: request.Email, Password: request.Password}
	err := ms.AuthenticationService.SignUp(user)
	if err != nil {
		return nil, err
	}
	return &pkg.SignUpResponse{Message: "Register has been successfully."}, nil
}
