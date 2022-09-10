package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type IJwtManager interface {
	GenerateJwtToken(payload interface{}) (token string, err error)
	VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error)
	RefreshJwtToken()
}

type JwtManager struct {
	expirationTime string
	secretKey      string
}

func NewJwtManager(expirationTime string, secretKey string) *JwtManager {
	return &JwtManager{expirationTime: expirationTime, secretKey: secretKey}
}

func (jm *JwtManager) GenerateJwtToken(payload interface{}) (token string, err error) {
	expiration, _ := strconv.Atoi(jm.expirationTime)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(time.Minute * time.Duration(expiration)).Unix(),
		"Subject":   payload,
	}).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error) {
	parsedToken, err := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"})).
		Parse(accessToken, func(t *jwt.Token) (any, error) {
			return []byte(jm.secretKey), nil
		})
	if err != nil {
		return nil, TokenSignature
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, TokenClaims
}
