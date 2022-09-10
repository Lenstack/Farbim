package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Masterminds/squirrel"
)

type IUserService interface {
	Show() (users []entities.User, err error)
	ShowBy(userId string) (user entities.User, err error)
	Update(userId string, newUser entities.User) (user entities.User, err error)
	Destroy(userId string) (rowsAffected int64, err error)
}

type UserService struct {
	userRepository repositories.UserRepository
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

func (us *UserService) Destroy(userId string) (rowsAffected int64, err error) {
	_, err = us.userRepository.GetUserById(userId)
	if err != nil {
		return 0, err
	}
	return us.userRepository.DestroyUser(userId)
}
