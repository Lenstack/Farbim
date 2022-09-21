package application

import (
	"errors"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignUp(_ context.Context, request *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {
	user := entities.User{Email: request.Email, Password: request.Password}
	if request.GetEmail() == "" || request.GetPassword() == "" {
		return nil, errors.New("email or password is empty")
	}

	err := ms.AuthenticationService.SignUp(user)
	if err != nil {
		return nil, err
	}
	return &pkg.SignUpResponse{Message: "Register has been successfully."}, nil
}
