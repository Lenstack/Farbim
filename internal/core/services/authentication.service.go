package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type IAuthenticationService interface {
	SignIn(user entities.User) (err error)
	SignUp(user entities.User) (err error)
	Logout(user entities.User) (err error)
}

type AuthenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(database *gorm.DB) *AuthenticationService {
	return &AuthenticationService{userRepository: repositories.UserRepository{
		Database: database,
	}}
}

func (as *AuthenticationService) SignIn(user entities.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (as *AuthenticationService) SignUp(user entities.User) (err error) {
	user.Id = uuid.New().String()
	user.Password = utils.HashPassword(user.Password)
	err = as.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) Logout(user entities.User) (err error) {
	//TODO implement me
	panic("implement me")
}
