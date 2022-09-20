package application

import (
	"errors"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

func (ms *MicroserviceServer) GetUser(ctx context.Context, _ *pkg.GetUserRequest) (*pkg.GetUserResponse, error) {
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

	user, err := ms.UserService.ShowBy(userId)
	if err != nil {
		return nil, err
	}

	userResponse := &pkg.GetUserResponse{
		Id:                     user.Id,
		Email:                  user.Email,
		AccessToken:            user.AccessToken,
		RefreshToken:           user.RefreshToken,
		LastResetSentAt:        user.LastResetSentAt,
		LastVerificationSentAt: user.LastVerificationSentAt,
		Verified:               user.Verified,
		CreatedAt:              timestamppb.New(user.CreatedAt),
		UpdatedAt:              timestamppb.New(user.UpdatedAt),
	}

	return userResponse, nil
}
