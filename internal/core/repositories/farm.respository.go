package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type IFarmRepository interface {
	CreateFarm(farm entities.Farm) (farmId string, err error)
}

type FarmRepository struct {
	Database squirrel.StatementBuilderType
}

func (fr *FarmRepository) CreateFarm(farm entities.Farm) (farmId string, err error) {
	farm.Id = uuid.New().String()
	bq := fr.Database.
		Insert(entities.FarmTableName).
		Columns("Id", "Name",
			"CategoryId", "UserId").
		Values(farm.Id, farm.Name, farm.CategoryId, farm.UserId).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&farmId)
	if err != nil {
		return "", utils.ErrorManager(err)
	}

	return farmId, nil
}
