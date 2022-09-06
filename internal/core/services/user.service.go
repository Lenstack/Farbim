package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"gorm.io/gorm"
)

type IUserService interface {
	Show() (users []entities.User, err error)
	ShowBy(userId int64) (user entities.User, err error)
	Update(userId int64, newUser entities.User) (err error)
	Destroy(userId int64) (err error)
}

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(database *gorm.DB) *UserService {
	return &UserService{
		userRepository: repositories.UserRepository{
			Database: database,
		},
	}
}

func (us *UserService) Show() (users []entities.User, err error) {
	return us.userRepository.Show()
}

func (us *UserService) ShowBy(userId int64) (user entities.User, err error) {
	return us.userRepository.ShowBy(userId)
}

func (us *UserService) Update(userId int64, newUser entities.User) (err error) {
	return us.userRepository.Update(userId, newUser)
}

func (us *UserService) Destroy(userId int64) (err error) {
	return us.userRepository.Destroy(userId)
}
