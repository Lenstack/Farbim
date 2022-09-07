package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUsers() (users []entities.User, err error)
	GetUserById(userId string) (user entities.User, err error)
	GetUserPasswordByEmail(email string) (password string, err error)
	GetUserIdByEmail(email string) (userId string, err error)
	CreateUser(user entities.User) (err error)
	UpdateUser(userId string, newUser entities.User) (err error)
	DestroyUser(userId string) (err error)
}

type UserRepository struct {
	Database *gorm.DB
}

func (ur *UserRepository) GetUsers() (users []entities.User, err error) {
	err = ur.Database.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUserById(userId string) (user entities.User, err error) {
	err = ur.Database.First(&user, "id", userId).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserPasswordByEmail(email string) (password string, err error) {
	user := entities.User{}
	err = ur.Database.Where("email", email).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

func (ur *UserRepository) GetUserIdByEmail(email string) (userId string, err error) {
	user := entities.User{}
	err = ur.Database.Where("email", email).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func (ur *UserRepository) CreateUser(user entities.User) (err error) {
	err = ur.Database.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) UpdateUser(userId string, newUser entities.User) (err error) {
	err = ur.Database.Model(&newUser).Where("id", userId).Updates(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DestroyUser(userId string) (err error) {
	err = ur.Database.Delete(&entities.User{}, "id", userId).Error
	if err != nil {
		return err
	}
	return nil
}
