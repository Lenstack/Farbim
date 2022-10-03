package application

import (
	"errors"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

func (ms *MicroserviceServer) Logout(ctx context.Context, _ *emptypb.Empty) (*pkg.LogoutResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	values := md.Get("Authorization")
	if len(values) == 0 {
		return nil, errors.New("missing authorization header")
	}

	headerToken := strings.Join(values, "")

	accessToken, err := ms.AuthenticationService.TokenManager.ExtractJwtToken(headerToken)
	if err != nil {
		return nil, err
	}

	_, err = ms.AuthenticationService.TokenManager.VerifyJwtToken(accessToken)
	if err != nil {
		return nil, err
	}

	err = ms.AuthenticationService.Logout("", accessToken)
	if err != nil {
		return nil, err
	}

	return &pkg.LogoutResponse{Message: "Logout has been successfully"}, nil
}
