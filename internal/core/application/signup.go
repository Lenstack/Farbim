package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignUp(ctx context.Context, request *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	userId, err := ms.AuthenticationService.SignUp(request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	return &pkg.SignUpResponse{Message: "Register has been successfully.", UserId: userId}, nil
}
