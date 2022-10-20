package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignIn(ctx context.Context, request *pkg.SignInRequest) (*pkg.SignInResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	accessToken, err := ms.AuthenticationService.SignIn(request.Email, request.Password, "")
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{AccessToken: accessToken}, nil
}
