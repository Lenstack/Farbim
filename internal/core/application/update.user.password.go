package application

import (
	"errors"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strings"
)

func (ms *MicroserviceServer) UpdateUserPassword(ctx context.Context, request *pkg.UpdateUserPasswordRequest) (*pkg.UpdateUserPasswordResponse, error) {
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

	if request.GetOldPassword() == "" || request.GetNewPassword() == "" {
		return nil, errors.New("oldPassword or newPassword is empty")
	}

	err := ms.UserService.UpdateUserPassword(userId, request.GetOldPassword(), request.GetNewPassword())
	if err != nil {
		return nil, err
	}
	return &pkg.UpdateUserPasswordResponse{Message: "your password has been updated successfully"}, nil
}
