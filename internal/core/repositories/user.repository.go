package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
)

type IUserRepository interface {
	GetUsers() (users []entities.User, err error)
	GetUserById(userId string) (user entities.User, err error)
	GetUserPasswordByEmail(email string) (password string, err error)
	GetUserIdByEmail(email string) (userId string, err error)
	CreateUser(user entities.User) (userId string, err error)
	UpdateUser(userId string, newUser entities.User) (user *entities.User, err error)
	DestroyUser(userId string) (rowsAffected int64, err error)
}

type UserRepository struct {
	Database squirrel.StatementBuilderType
}

func (ur *UserRepository) GetUsers() (users []entities.User, err error) {
	user := entities.User{}
	bq := ur.Database.Select("Id", "Email",
		"Token", "LastResetSentAt", "LastVerificationSentAt",
		"Verified", "CreatedAt", "UpdatedAt").From(entities.UserTableName)

	rows, err := bq.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Token,
			&user.LastResetSentAt, &user.LastVerificationSentAt,
			&user.Verified, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) GetUserById(userId string) (user entities.User, err error) {
	bq := ur.Database.Select("Id", "Email",
		"Token", "LastResetSentAt", "LastVerificationSentAt",
		"Verified", "CreatedAt", "UpdatedAt").
		From(entities.UserTableName).
		Where(squirrel.Eq{"id": userId})

	err = bq.QueryRow().Scan(&user.Id, &user.Email, &user.Token,
		&user.LastResetSentAt, &user.LastVerificationSentAt,
		&user.Verified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserPasswordByEmail(email string) (password string, err error) {
	bq := ur.Database.Select("password").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (ur *UserRepository) GetUserIdByEmail(email string) (userId string, err error) {
	bq := ur.Database.Select("Id").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (ur *UserRepository) CreateUser(user entities.User) (userId string, err error) {
	bq := ur.Database.
		Insert(entities.UserTableName).
		Columns("Id", "Email", "Password",
			"Token", "LastResetSentAt", "LastVerificationSentAt",
			"Verified", "CreatedAt", "UpdatedAt").
		Values(user.Id, user.Email, user.Password, user.Token,
			user.LastResetSentAt, user.LastVerificationSentAt,
			user.Verified, user.CreatedAt, user.UpdatedAt).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (ur *UserRepository) UpdateUser(userId string, newUser entities.User) (user entities.User, err error) {
	userMap := map[string]interface{}{"email": newUser.Email, "password": newUser.Password}
	bq := ur.Database.
		Update(entities.UserTableName).
		SetMap(userMap).
		Where(squirrel.Eq{"id": userId}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&user.Id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) DestroyUser(userId string) (rowsAffected int64, err error) {
	bq := ur.Database.Delete(entities.UserTableName).Where(squirrel.Eq{"id": userId})
	result, err := bq.Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
