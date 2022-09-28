package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v9"
)

type IAuthenticationService interface {
	SignIn(user entities.User) (accessToken string, err error)
	SignUp(user entities.User) (err error)
	Logout(userId string, token string) (err error)
	VerifyAccount(userId string) (err error)
	DisableAccount(userId string) (err error)
}

type AuthenticationService struct {
	userRepository  repositories.UserRepository
	redisRepository repositories.RedisRepository
	TokenManager    utils.JwtManager
	bcryptManager   utils.BcryptManager
	redisManager    *redis.Client
}

func NewAuthenticationService(database squirrel.StatementBuilderType, tokenManager utils.JwtManager, redisManager *redis.Client) *AuthenticationService {
	return &AuthenticationService{
		userRepository: repositories.UserRepository{
			Database: database,
		},
		redisRepository: repositories.RedisRepository{
			RedisManager: redisManager,
		},
		TokenManager: tokenManager,
	}
}

func (as *AuthenticationService) SignIn(user entities.User) (accessToken string, err error) {
	verified, err := as.userRepository.GetUserVerifiedByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !verified {
		return "", utils.EmailIsNotVerified
	}

	hashedPassword, err := as.userRepository.GetUserPasswordByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = as.bcryptManager.CompareHashedPassword(hashedPassword, user.Password)
	if err != nil {
		return "", utils.ErrorManager(err)
	}

	userId, err := as.userRepository.GetUserIdByEmail(user.Email)
	if err != nil {
		return "", err
	}

	roles, err := as.userRepository.GetUserRolesById(userId)
	if err != nil {
		return "", err
	}

	accessToken, err = as.TokenManager.GenerateJwtAccessToken(utils.PayloadClaims{Id: userId, Roles: []string{roles}})
	if err != nil {
		return "", err
	}

	_, err = as.userRepository.UpdateUserAccessToken(userId, accessToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (as *AuthenticationService) SignUp(user entities.User) (err error) {
	idUserExist, err := as.userRepository.GetUserIdByEmail(user.Email)
	if idUserExist != "" {
		return utils.ItemAlreadyExist
	}

	userId, err := as.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	_, err = as.userRepository.CreateProfile(entities.Profile{UserId: userId})
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) Logout(userId string, token string) (err error) {
	err = as.redisRepository.IsTokenBlacklisted(token)
	if err != nil {
		return err
	}
	err = as.redisRepository.AddTokenToBlacklist(token)
	if err != nil {
		return err
	}
	_, err = as.userRepository.UpdateUserAccessToken(userId, "")
	if err != nil {
		return err
	}
	_, err = as.userRepository.UpdateUserRefreshToken(userId, "")
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) VerifyAccount(userId string) (err error) {
	_, err = as.userRepository.UpdateUserVerifiedAccountById(userId)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthenticationService) DisableAccount(userId string) (err error) {
	_, err = as.userRepository.UpdateUserDisableAccountById(userId)
	if err != nil {
		return err
	}
	return nil
}
