package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (ms *MicroserviceServer) GetUsers(_ context.Context, _ *emptypb.Empty) (*pkg.GetUsersResponse, error) {
	users, err := ms.UserService.Show()
	if err != nil {
		return nil, err
	}

	usersResponse := make([]*pkg.GetUserResponse, 0, len(users))
	for _, user := range users {
		usersResponse = append(usersResponse, &pkg.GetUserResponse{
			Id:                     user.Id,
			Email:                  user.Email,
			AccessToken:            user.AccessToken,
			RefreshToken:           user.RefreshToken,
			LastResetSentAt:        user.LastResetSentAt,
			LastVerificationSentAt: user.LastVerificationSentAt,
			Verified:               user.Verified,
			CreatedAt:              timestamppb.New(user.CreatedAt),
			UpdatedAt:              timestamppb.New(user.UpdatedAt),
		})
	}

	return &pkg.GetUsersResponse{Page: 1, PerPage: 10, Total: int32(len(usersResponse)), Users: usersResponse}, nil
}
