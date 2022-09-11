package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
)

type IAuthenticationService interface {
	SignIn(user entities.User) (accessToken string, refreshToken string, err error)
	SignUp(user entities.User) (err error)
	Logout(userId string) (err error)
}

type AuthenticationService struct {
	userRepository repositories.UserRepository
	tokenManager   utils.JwtManager
	bcryptManager  utils.BcryptManager
}

func NewAuthenticationService(database squirrel.StatementBuilderType, tokenManager utils.JwtManager) *AuthenticationService {
	return &AuthenticationService{
		userRepository: repositories.UserRepository{
			Database: database,
		},
		tokenManager: tokenManager,
	}
}

func (as *AuthenticationService) SignIn(user entities.User) (accessToken string, refreshToken string, err error) {
	hashedPassword, err := as.userRepository.GetUserPasswordByEmail(user.Email)
	if err != nil {
		return "", "", err
	}

	err = as.bcryptManager.CompareHashedPassword(hashedPassword, user.Password)
	if err != nil {
		return "", "", utils.ErrorManager(err)
	}

	userId, err := as.userRepository.GetUserIdByEmail(user.Email)
	if err != nil {
		return "", "", err
	}

	accessToken, err = as.tokenManager.GenerateJwtAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = as.tokenManager.GenerateJwtAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	_, err = as.userRepository.UpdateUserTokens(userId, accessToken, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (as *AuthenticationService) SignUp(user entities.User) (err error) {
	idUserExist, err := as.userRepository.GetUserIdByEmail(user.Email)
	if idUserExist != "" {
		return utils.ItemAlreadyExist
	}

	_, err = as.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthenticationService) Logout(userId string) (err error) {
	//TODO implement me
	panic("implement me")
}
