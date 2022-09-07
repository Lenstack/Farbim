package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type IJwtManager interface {
	GenerateJwtToken(payload interface{}) (string, error)
	ValidateJwtToken(accessToken string) (interface{}, error)
}

type JwtManager struct {
	expirationTime string
	secretKey      string
}

func NewJwtManager(expirationTime string, secretKey string) *JwtManager {
	return &JwtManager{expirationTime: expirationTime, secretKey: secretKey}
}

func (jm *JwtManager) GenerateJwtToken(payload interface{}) (string, error) {
	expiration, _ := strconv.Atoi(jm.expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(time.Minute * time.Duration(expiration)).Unix(),
		"Subject":   payload,
	})
	return token.SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) ValidateJwtToken(accessToken string) (interface{}, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected Signing Method ")
		}
		return []byte(jm.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Cannot Get Claims From Token ")
	}
	return claims["Subject"], nil
}
