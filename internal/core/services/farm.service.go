package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Masterminds/squirrel"
)

type IFarmService interface {
	CreateFarm(farm entities.Farm) (farmId string, err error)
}

type FarmService struct {
	farmRepository repositories.FarmRepository
}

func NewFarmService(database squirrel.StatementBuilderType) *FarmService {
	return &FarmService{
		farmRepository: repositories.FarmRepository{
			Database: database,
		},
	}
}

func (fs *FarmService) CreateFarm(farm entities.Farm) (farmId string, err error) {
	return fs.farmRepository.CreateFarm(farm)
}
