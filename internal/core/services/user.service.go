package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
)

type IUserService interface {
	Show() (users []entities.User, err error)
	ShowBy(userId string) (user entities.User, err error)
	Update(userId string, newUser entities.User) (user entities.User, err error)
	UpdateUserRefreshToken(userId string, refreshToken string) (user entities.User, err error)
	UpdateUserPassword(userId string, oldPassword string, newPassword string) (err error)
	Destroy(userId string) (rowsAffected int64, err error)
}

type UserService struct {
	userRepository repositories.UserRepository
	bcryptManager  utils.BcryptManager
}

func NewUserService(database squirrel.StatementBuilderType) *UserService {
	return &UserService{
		userRepository: repositories.UserRepository{
			Database: database,
		},
	}
}

func (us *UserService) Show() (users []entities.User, err error) {
	return us.userRepository.GetUsers()
}

func (us *UserService) ShowBy(userId string) (user entities.User, err error) {
	return us.userRepository.GetUserById(userId)
}

func (us *UserService) Update(userId string, newUser entities.User) (user entities.User, err error) {
	_, err = us.userRepository.GetUserById(userId)
	if err != nil {
		return entities.User{}, err
	}
	return us.userRepository.UpdateUser(userId, newUser)
}

func (us *UserService) UpdateUserRefreshToken(userId string, refreshToken string) (user entities.User, err error) {
	return us.userRepository.UpdateUserRefreshToken(userId, refreshToken)
}

func (us *UserService) Destroy(userId string) (rowsAffected int64, err error) {
	_, err = us.userRepository.GetUserById(userId)
	if err != nil {
		return 0, err
	}
	return us.userRepository.DestroyUser(userId)
}

func (us *UserService) UpdateUserPassword(userId string, oldPassword string, newPassword string) (err error) {
	hashedOldPassword, err := us.userRepository.GetUserPasswordById(userId)
	if err != nil {
		return err
	}

	err = us.bcryptManager.CompareHashedPassword(hashedOldPassword, oldPassword)
	if err != nil {
		return err
	}

	hashedNewPassword := us.bcryptManager.HashPassword(newPassword)
	_, err = us.userRepository.UpdateUserPassword(userId, hashedNewPassword)
	if err != nil {
		return err
	}

	return nil
}
