package services

import (
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

type IAuthenticationService interface {
	SignIn(email string, password string) (accessToken string, err error)
	SignUp(email string, password string) (userId string, err error)
	Logout(headerToken string) (userId string, err error)
	VerifyEmail(headerToken string) (userId string, err error)
}

type AuthenticationService struct {
	userRepository    repositories.UserRepository
	profileRepository repositories.ProfileRepository
	sessionRepository repositories.SessionRepository
	tokenManager      utils.JwtManager
	bcryptManager     utils.BcryptManager
	emailManager      utils.EmailManger
	redisManager      *redis.Client
}

func NewAuthenticationService(database squirrel.StatementBuilderType, tokenManager utils.JwtManager, emailManager utils.EmailManger) *AuthenticationService {
	return &AuthenticationService{
		userRepository: repositories.UserRepository{
			Database: database,
		},
		profileRepository: repositories.ProfileRepository{
			Database: database,
		},
		sessionRepository: repositories.SessionRepository{
			Database: database,
		},
		tokenManager: tokenManager,
		emailManager: emailManager,
	}
}

func (as *AuthenticationService) SignIn(email string, password string, location string) (accessToken string, err error) {
	usr, err := as.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", status.Errorf(codes.NotFound, "this account not exist")
	}

	if !usr.Verified {
		return "", status.Errorf(codes.Unauthenticated, "this account has not been verified")
	}

	err = as.bcryptManager.CompareHashedPassword(usr.Password, password)
	if err != nil {
		return "", status.Errorf(codes.Internal, "email or password are wrong")
	}

	accessToken, err = as.tokenManager.GenerateJwtAccessToken(usr.Id, viper.Get("JWT_EXPIRATION_ACCESS").(string))
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}

	expirationAccess, _ := strconv.Atoi(viper.Get("JWT_EXPIRATION_ACCESS").(string))
	_, err = as.sessionRepository.CreateSession(entities.Session{
		Id:           uuid.New().String(),
		Type:         "Access",
		Token:        accessToken,
		ClientOrigin: location,
		ExpiredAt:    time.Now().Add(time.Minute * time.Duration(expirationAccess)).Format("2006-01-02 15:04:05.000000"),
		UserId:       usr.Id,
	})
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}

	return accessToken, err
}

func (as *AuthenticationService) SignUp(email string, password string) (userId string, err error) {
	isUserExist, err := as.userRepository.GetUserIdByEmail(email)
	if isUserExist != "" {
		return "", status.Errorf(codes.Internal, "this email is already registered")
	}

	hashedPassword := as.bcryptManager.HashPassword(password)
	user := entities.User{
		Id:              uuid.New().String(),
		Email:           email,
		Password:        hashedPassword,
		TokenKey:        uuid.New().String(),
		RolesId:         []string{"User"},
		LastResetSentAt: time.Now().Format("2006-01-02 15:04:05.000000"),
	}
	userId, err = as.userRepository.CreateUser(user)
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}

	profile := entities.Profile{
		Id:     uuid.New().String(),
		Avatar: "1",
		UserId: userId,
	}
	_, err = as.profileRepository.CreateProfile(profile)
	if err != nil {
		return "", err
	}

	go func() {
		verifyEmailLink := fmt.Sprintf("Hello %s\n\nPlease verify your account by clicking the next link:\nhttp://%s/v1/authentication/verify_email/%s\n\nThank You!", user.Email, "localhost:8080", user.TokenKey)
		err = as.emailManager.SendEmail(user.Email, "Account Verification Link", verifyEmailLink)
		if err != nil {
			return
		}
	}()

	return userId, err
}

func (as *AuthenticationService) Logout(headerToken string) (userId string, err error) {
	return userId, nil
}

func (as *AuthenticationService) VerifyEmail(headerToken string) (userId string, err error) {
	id, err := as.userRepository.GetUserByTokenKey(headerToken)
	if err != nil {
		return "", status.Errorf(codes.Internal, "your verification link may have expired")
	}

	userId, err = as.userRepository.UpdateVerifiedByUserId(id, true)
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}
	_, err = as.userRepository.UpdateTokenKeyByUserId(userId)
	if err != nil {
		return "", status.Errorf(codes.Internal, err.Error())
	}
	return userId, nil
}
