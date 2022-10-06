package application

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignIn(_ context.Context, request *pkg.SignInRequest) (*pkg.SignInResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user := entities.User{Email: request.Email, Password: request.Password}
	token, err := ms.AuthenticationService.SignIn(user)
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{Token: token}, nil
}
