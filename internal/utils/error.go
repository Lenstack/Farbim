package utils

import (
	"errors"
	"fmt"
)

var (
	AccessNotAuthorizated   = errors.New("access not authorized")
	ItemNotExist            = errors.New("item not exist")
	ItemWithout             = errors.New("without items")
	ItemAlreadyExist        = errors.New("the item already exist")
	AuthenticationIncorrect = errors.New("email or password are incorrect")
	TokenIncorrect          = errors.New("the token is incorrect")
	TokenInvalid            = errors.New("the token is invalid")
	TokenWithout            = errors.New("without token")
)

func ErrorManager(err error) error {
	fmt.Println("error manager:", err)

	switch err.Error() {
	case "sql: no rows in result set":
		return ItemNotExist
	case "pq: llave duplicada viola restricción de unicidad «users_email_key»":
		return ItemAlreadyExist
	case "crypto/bcrypt: hashedPassword is not the hash of the given password":
		return AuthenticationIncorrect
	default:
		return err
	}
}
