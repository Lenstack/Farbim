package repositories

import (
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type IUserRepository interface {
	GetUserById(userId string) (user entities.User, err error)
	GetUserByEmail(email string) (user entities.User, err error)
	GetUserIdByEmail(email string) (userId string, err error)
	GetUserPasswordByEmail(email string) (password string, err error)
	GetUserByTokenKey(TokenKey string) (userId string, err error)
	GetUserRolesId(userId string) (rolesId []string, err error)
	CreateUser(user entities.User) (userId string, err error)
	UpdateVerifiedByUserId(id string, isVerified bool) (userId string, err error)
	UpdateTokenKeyByUserId(id string) (userId string, err error)
	UpdateRolesIdByUserId(id string, rolesId []string) (userId string, err error)
}

type UserRepository struct {
	Database squirrel.StatementBuilderType
}

func (ur *UserRepository) GetUserById(userId string) (user entities.User, err error) {
	bq := ur.Database.Select("Id", "Email", "Password", "Verified", "RolesId", "Code", "CreatedAt", "UpdatedAt").
		From(entities.UserTableName).
		Where(squirrel.Eq{"id": userId})

	err = bq.QueryRow().Scan(&user.Id, &user.Email, &user.Password, &user.Verified,
		&user.RolesId, &user.Code, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (user entities.User, err error) {
	bq := ur.Database.Select("Id", "Email", "Password", "Verified", "CreatedAt", "UpdatedAt").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&user.Id, &user.Email, &user.Password, &user.Verified,
		&user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return entities.User{}, err
	}

	return user, nil
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

func (ur *UserRepository) GetUserPasswordByEmail(email string) (password string, err error) {
	bq := ur.Database.Select("Password").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (ur *UserRepository) GetUserByTokenKey(TokenKey string) (userId string, err error) {
	bq := ur.Database.Select("id").
		From(entities.UserTableName).
		Where(squirrel.Eq{"tokenkey": TokenKey})

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (ur *UserRepository) GetUserRolesId(userId string) (rolesId []string, err error) {
	bq := ur.Database.Select("RolesId").
		From(entities.UserTableName).
		Where(squirrel.Eq{"id": userId})

	err = bq.QueryRow().Scan(pq.Array(&rolesId))
	if err != nil {
		return nil, err
	}

	return rolesId, nil
}

func (ur *UserRepository) CreateUser(user entities.User) (userId string, err error) {
	bq := ur.Database.
		Insert(entities.UserTableName).
		Columns("Id", "Email", "Password", "TokenKey", "RolesId", "LastResetSentAt").
		Values(user.Id, user.Email, user.Password, user.TokenKey, pq.Array(user.RolesId), user.LastResetSentAt).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (ur *UserRepository) UpdateVerifiedByUserId(id string, isVerified bool) (userId string, err error) {
	bq := ur.Database.
		Update(entities.UserTableName).
		Set("verified", isVerified).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (ur *UserRepository) UpdateTokenKeyByUserId(id string) (userId string, err error) {
	bq := ur.Database.
		Update(entities.UserTableName).
		Set("tokenkey", nil).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (ur *UserRepository) UpdateRolesIdByUserId(id string, rolesId []string) (userId string, err error) {
	bq := ur.Database.
		Update(entities.UserTableName).
		Set("RolesId", pq.Array(rolesId)).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
