package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetUsers() (users []entities.User, err error)
	GetUserById(userId string) (user entities.User, err error)
	GetUserPasswordByEmail(email string) (password string, err error)
	GetUserIdByEmail(email string) (userId string, err error)
	GetUserVerifiedByEmail(email string) (verified bool, err error)
	CreateUser(user entities.User) (userId string, err error)
	UpdateUser(userId string, newUser entities.User) (user entities.User, err error)
	UpdateUserAccessToken(userId string, accessToken string) (user entities.User, err error)
	UpdateUserRefreshToken(userId string, refreshToken string) (user entities.User, err error)
	DestroyUser(userId string) (rowsAffected int64, err error)
}

type UserRepository struct {
	Database      squirrel.StatementBuilderType
	BcryptManager utils.BcryptManager
}

func (ur *UserRepository) GetUsers() (users []entities.User, err error) {
	user := entities.User{}
	bq := ur.Database.Select("Id", "Email",
		"AccessToken", "RefreshToken", "LastResetSentAt", "LastVerificationSentAt",
		"Verified", "CreatedAt", "UpdatedAt").From(entities.UserTableName)

	rows, err := bq.Query()
	if err != nil {
		return nil, utils.ErrorManager(err)
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email,
			&user.AccessToken, &user.RefreshToken,
			&user.LastResetSentAt, &user.LastVerificationSentAt,
			&user.Verified, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, utils.ItemWithout
	}

	return users, nil
}

func (ur *UserRepository) GetUserById(userId string) (user entities.User, err error) {
	bq := ur.Database.Select("Id", "Email",
		"AccessToken", "RefreshToken", "LastResetSentAt", "LastVerificationSentAt",
		"Verified", "CreatedAt", "UpdatedAt").
		From(entities.UserTableName).
		Where(squirrel.Eq{"id": userId})

	err = bq.QueryRow().Scan(&user.Id, &user.Email,
		&user.AccessToken, &user.RefreshToken,
		&user.LastResetSentAt, &user.LastVerificationSentAt,
		&user.Verified, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return entities.User{}, utils.ErrorManager(err)
	}
	return user, nil
}

func (ur *UserRepository) GetUserPasswordByEmail(email string) (password string, err error) {
	bq := ur.Database.Select("password").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&password)
	if err != nil {
		return "", utils.AuthenticationIncorrect
	}
	return password, nil
}

func (ur *UserRepository) GetUserIdByEmail(email string) (userId string, err error) {
	bq := ur.Database.Select("Id").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", utils.ErrorManager(err)
	}
	return userId, nil
}

func (ur *UserRepository) GetUserVerifiedByEmail(email string) (verified bool, err error) {
	bq := ur.Database.Select("Verified").
		From(entities.UserTableName).
		Where(squirrel.Eq{"email": email})

	err = bq.QueryRow().Scan(&verified)
	if err != nil {
		return false, utils.ErrorManager(err)
	}
	return verified, nil
}

func (ur *UserRepository) CreateUser(user entities.User) (userId string, err error) {
	user.Id = uuid.New().String()
	user.Password = ur.BcryptManager.HashPassword(user.Password)

	bq := ur.Database.
		Insert(entities.UserTableName).
		Columns("Id", "Email", "Password",
			"AccessToken", "RefreshToken",
			"LastResetSentAt", "LastVerificationSentAt",
			"Verified").
		Values(user.Id, user.Email, user.Password,
			user.AccessToken, user.RefreshToken,
			user.LastResetSentAt, user.LastVerificationSentAt,
			user.Verified).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&userId)
	if err != nil {
		return "", utils.ErrorManager(err)
	}

	return userId, nil
}

func (ur *UserRepository) UpdateUser(userId string, newUser entities.User) (user entities.User, err error) {
	userMap := map[string]interface{}{"email": newUser.Email, "password": ur.BcryptManager.HashPassword(newUser.Password)}
	bq := ur.Database.
		Update(entities.UserTableName).
		SetMap(userMap).
		Where(squirrel.Eq{"id": userId}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&user.Id)
	if err != nil {
		return entities.User{}, utils.ErrorManager(err)
	}

	return user, nil
}

func (ur *UserRepository) UpdateUserAccessToken(userId string, accessToken string) (user entities.User, err error) {
	userMap := map[string]interface{}{"accessToken": accessToken}
	bq := ur.Database.
		Update(entities.UserTableName).
		SetMap(userMap).
		Where(squirrel.Eq{"id": userId}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&user.Id)
	if err != nil {
		return entities.User{}, utils.ErrorManager(err)
	}
	return user, nil
}

func (ur *UserRepository) UpdateUserRefreshToken(userId string, refreshToken string) (user entities.User, err error) {
	userMap := map[string]interface{}{"refreshToken": refreshToken}
	bq := ur.Database.
		Update(entities.UserTableName).
		SetMap(userMap).
		Where(squirrel.Eq{"id": userId}).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&user.Id)
	if err != nil {
		return entities.User{}, utils.ErrorManager(err)
	}
	return user, nil
}

func (ur *UserRepository) DestroyUser(userId string) (rowsAffected int64, err error) {
	bq := ur.Database.Delete(entities.UserTableName).Where(squirrel.Eq{"id": userId})
	result, err := bq.Exec()
	if err != nil {
		return 0, utils.ErrorManager(err)
	}
	return result.RowsAffected()
}
