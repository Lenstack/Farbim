package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

type IJwtManager interface {
	GenerateJwtAccessToken(payload interface{}) (token string, err error)
	GenerateJwtRefreshToken(payload interface{}) (token string, err error)
	VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error)
	Test() (string, error)
}

type JwtManager struct {
	expirationToken   string
	expirationRefresh string
	secretKey         string
}

func NewJwtManager(expirationToken string, expirationRefresh string, secretKey string) *JwtManager {
	return &JwtManager{expirationToken: expirationToken, expirationRefresh: expirationRefresh, secretKey: secretKey}
}

func (jm *JwtManager) GenerateJwtAccessToken(payload interface{}) (accessToken string, err error) {
	expirationAccessToken, _ := strconv.Atoi(jm.expirationToken)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid.New().String(),
		"exp":  time.Now().Add(time.Minute * time.Duration(expirationAccessToken)).Unix(),
		"sub":  payload,
	}).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) GenerateJwtRefreshToken(payload interface{}) (refreshToken string, err error) {
	expirationRefreshToken, _ := strconv.Atoi(jm.expirationRefresh)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid.New().String(),
		"exp":  time.Now().Add(time.Hour * 24 * time.Duration(expirationRefreshToken)).Unix(),
		"sub":  payload,
	}).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error) {
	parsedToken, err := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"})).
		Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(jm.secretKey), nil
		})
	if err != nil {
		return nil, ErrorManager(err)
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, TokenClaims
}

func (jm *JwtManager) Test() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
