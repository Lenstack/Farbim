package application

import (
	"errors"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

func (ms *MicroserviceServer) VerifyAccount(ctx context.Context, _ *emptypb.Empty) (*pkg.VerifyAccountResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	values := md.Get("Id")
	if len(values) == 0 {
		return nil, errors.New("missing metadata key Id")
	}

	userId := strings.ReplaceAll(values[0], " ", "")
	if userId == "" {
		return nil, errors.New("the value of metadata key Id this empty")
	}
	err := ms.AuthenticationService.VerifyAccount(userId)
	if err != nil {
		return nil, err
	}
	return &pkg.VerifyAccountResponse{Message: "the account has been activated"}, nil
}
