package application

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignIn(_ context.Context, req *pkg.SignInRequest) (*pkg.SignInResponse, error) {
	user := entities.User{Email: req.Email, Password: req.Password}
	token, err := ms.AuthenticationService.SignIn(user)
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{Token: token}, nil
}
