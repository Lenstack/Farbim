package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
)

type IProfileRepository interface {
	UpdateProfile(userId string, newProfile entities.Profile) (userIdAffected string, err error)
}

type ProfileRepository struct {
	Database squirrel.StatementBuilderType
}

func (pr *ProfileRepository) UpdateProfile(userId string, newProfile entities.Profile) (userIdAffected string, err error) {
	userMap := map[string]interface{}{"name": newProfile.Name, "avatar": newProfile.Avatar}
	bq := pr.Database.
		Update(entities.ProfileTableName).
		SetMap(userMap).
		Where(squirrel.Eq{"userid": userId}).
		Suffix("RETURNING id")

	err = bq.QueryRow().Scan(&userIdAffected)
	if err != nil {
		return "", err
	}
	return userIdAffected, nil
}
