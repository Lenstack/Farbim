package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (ms *MicroserviceServer) GetUser(_ context.Context, request *pkg.GetUserRequest) (*pkg.GetUserResponse, error) {
	user, err := ms.UserService.ShowBy(request.Id)
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
