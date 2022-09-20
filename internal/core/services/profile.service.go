package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Masterminds/squirrel"
)

type IProfileService interface {
	UpdateProfile(userId string, newProfile entities.Profile) (userIdAffected string, err error)
}

type ProfileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(database squirrel.StatementBuilderType) *ProfileService {
	return &ProfileService{
		profileRepository: repositories.ProfileRepository{
			Database: database,
		},
	}
}

func (ps *ProfileService) UpdateProfile(userId string, newProfile entities.Profile) (userIdAffected string, err error) {
	return ps.profileRepository.UpdateProfile(userId, newProfile)
}
