package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"gorm.io/gorm"
)

type IUserService interface {
	List() (users []entities.User, err error)
	ShowBy(userId int64) (user entities.User, err error)
	Create(user entities.User) (err error)
	Update(userId int64, newUser entities.User) (err error)
	Destroy(userId int64) (err error)
	Verify(user entities.User) (err error)
}

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(database *gorm.DB) *UserService {
	return &UserService{
		UserRepository: repositories.UserRepository{
			Database: database,
		},
	}
}

func (us *UserService) List() (users []entities.User, err error) {
	return us.UserRepository.List()
}

func (us *UserService) ShowBy(userId int64) (user entities.User, err error) {
	return us.UserRepository.ShowBy(userId)
}

func (us *UserService) Create(user entities.User) (err error) {
	return us.UserRepository.Create(user)
}

func (us *UserService) Update(userId int64, newUser entities.User) (err error) {
	return us.UserRepository.Update(userId, newUser)
}

func (us *UserService) Destroy(userId int64) (err error) {
	return us.UserRepository.Destroy(userId)
}

func (us *UserService) Verify(user entities.User) (err error) {
	return us.UserRepository.Verify(user)
}
