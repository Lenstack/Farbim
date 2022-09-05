package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	List() (users []entities.User, err error)
	ShowBy(userId int64) (user entities.User, err error)
	Create(user entities.User) (err error)
	Update(userId int64, newUser entities.User) (err error)
	Destroy(userId int64) (err error)
	Verify(user entities.User) (err error)
}

type UserRepository struct {
	Database *gorm.DB
}

func (ur *UserRepository) List() (users []entities.User, err error) {
	results := ur.Database.Find(&users)
	if results.Error != nil {
		return nil, results.Error
	}
	return users, nil
}

func (ur *UserRepository) ShowBy(userId int64) (user entities.User, err error) {
	result := ur.Database.First(&user, &userId)
	if result.Error != nil {

	}
	return user, nil
}

func (ur *UserRepository) Create(user entities.User) (err error) {
	result := ur.Database.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *UserRepository) Update(userId int64, newUser entities.User) (err error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) Destroy(userId int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) Verify(user entities.User) (err error) {
	//TODO implement me
	panic("implement me")
}
