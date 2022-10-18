package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
)

func (ms *MicroserviceServer) SignIn(ctx context.Context, request *pkg.SignInRequest) (*pkg.SignInResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	p, _ := peer.FromContext(ctx)
	accessToken, err := ms.AuthenticationService.SignIn(request.Email, request.Password, p.Addr.Network()+p.Addr.String())
	if err != nil {
		return nil, err
	}
	return &pkg.SignInResponse{AccessToken: accessToken}, nil
}
