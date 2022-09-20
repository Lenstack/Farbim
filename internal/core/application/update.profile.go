package application

import (
	"errors"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strings"
)

func (ms *MicroserviceServer) UpdateProfile(ctx context.Context, request *pkg.UpdateProfileRequest) (*pkg.UpdateProfileResponse, error) {
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

	if request.GetName() == "" || request.GetAvatar() == "" {
		return nil, errors.New("name or avatar is empty")
	}

	profile := entities.Profile{Name: request.GetName(), Avatar: request.GetAvatar()}
	_, err := ms.ProfileService.UpdateProfile(userId, profile)
	if err != nil {
		return nil, err
	}
	return &pkg.UpdateProfileResponse{Message: "the profile has been updated successfully"}, err
}
