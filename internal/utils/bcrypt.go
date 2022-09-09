package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type IBcryptManager interface {
	HashPassword(password string) string
	CompareHashedPassword(hashedPassword, password string) error
}

type BcryptManager struct {
}

func (bm *BcryptManager) HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func (bm *BcryptManager) CompareHashedPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
