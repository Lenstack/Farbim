package utils

import (
	"errors"
)

var (
	AccessNotAuthorizated   = errors.New("access not authorized")
	EmailIsNotVerified      = errors.New("email is not verified")
	ItemNotExist            = errors.New("item not exist")
	ItemWithout             = errors.New("without items")
	ItemAlreadyExist        = errors.New("the item already exist")
	AuthenticationIncorrect = errors.New("email or password are incorrect")
	TokenExpired            = errors.New("the token has been expired")
	TokenIncorrect          = errors.New("the token is incorrect")
	TokenInvalid            = errors.New("the token is invalid or malformed")
	TokenWithout            = errors.New("authentication header no contains bearer token")
	TokenSignature          = errors.New("the token signature is invalid")
	TokenClaims             = errors.New("cannot get claims from token")
)

func ErrorManager(err error) error {
	switch err.Error() {
	case "sql: no rows in result set":
		return ItemNotExist
	case "pq: llave duplicada viola restricción de unicidad «users_email_key»":
		return ItemAlreadyExist
	case "crypto/bcrypt: hashedPassword is not the hash of the given password":
		return AuthenticationIncorrect
	case "Token is expired":
		return TokenExpired
	case "illegal base64 data at input byte 36":
		return TokenIncorrect
	case "signature is invalid":
		return TokenSignature
	case "token contains an invalid number of segments":
		return TokenInvalid
	default:
		return err
	}
}
