package application

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) CreateFarm(_ context.Context, request *pkg.CreateFarmRequest) (*pkg.CreateFarmResponse, error) {
	farm := entities.Farm{Name: request.GetName(), CategoryId: request.CategoryId, UserId: request.GetUserId()}
	_, err := ms.CategoryService.GetCategory(farm.CategoryId)
	if err != nil {
		return nil, err
	}

	_, err = ms.FarmService.CreateFarm(farm)
	if err != nil {
		return nil, err
	}
	return &pkg.CreateFarmResponse{Message: "the farm has been created successfully."}, err
}
