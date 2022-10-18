package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
)

type IProfileRepository interface {
	CreateProfile(profile entities.Profile) (profileId string, err error)
}

type ProfileRepository struct {
	Database squirrel.StatementBuilderType
}

func (pr *ProfileRepository) CreateProfile(profile entities.Profile) (profileId string, err error) {
	bq := pr.Database.
		Insert(entities.ProfileTableName).
		Columns("Id", "Name", "UserId", "Avatar").
		Values(profile.Id, profile.Name, profile.UserId, profile.Avatar).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&profileId)
	if err != nil {
		return "", err
	}
	return profileId, nil
}
