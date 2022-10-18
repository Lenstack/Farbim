package application

import (
	"errors"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strings"
)

func (ms *MicroserviceServer) Logout(ctx context.Context, _ *pkg.LogoutRequest) (*pkg.LogoutResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	values := md.Get("Authorization")
	if len(values) == 0 {
		return nil, errors.New("missing authorization header")
	}

	headerToken := strings.Join(values, "")

	_, err := ms.AuthenticationService.Logout(headerToken)
	if err != nil {
		return nil, err
	}

	return &pkg.LogoutResponse{Message: "Logout has been successfully"}, nil
}
