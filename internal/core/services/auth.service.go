package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type IAuthenticationService interface {
	SignIn(user entities.User) (token string, err error)
	SignUp(user entities.User) (token string, err error)
	Logout(userId string) (err error)
}

type AuthenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(database squirrel.StatementBuilderType) *AuthenticationService {
	return &AuthenticationService{userRepository: repositories.UserRepository{
		Database: database,
	}}
}

func (as *AuthenticationService) SignIn(user entities.User) (token string, err error) {
	hashedPassword, err := as.userRepository.GetUserPasswordByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = utils.CompareHashedPassword(hashedPassword, user.Password)
	if err != nil {
		return "", err
	}

	userId, err := as.userRepository.GetUserIdByEmail(user.Email)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (as *AuthenticationService) SignUp(user entities.User) (token string, err error) {
	user.Id = uuid.New().String()
	user.Password = utils.HashPassword(user.Password)

	err = as.userRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (as *AuthenticationService) Logout(user entities.User) (err error) {
	//TODO implement me
	panic("implement me")
}
