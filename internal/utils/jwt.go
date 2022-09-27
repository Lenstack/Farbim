package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type IJwtManager interface {
	GenerateJwtAccessToken(payload interface{}) (accessToken string, err error)
	GenerateJwtRefreshToken(payload interface{}) (refreshToken string, err error)
	VerifyJwtToken(accessToken string) (claims string, err error)
	ExtractJwtToken(headerToken string) (clearedToken string, err error)
}

type JwtManager struct {
	expirationToken   string
	expirationRefresh string
	secretKey         string
}

type PayloadClaims struct {
	UserId string
	Roles  string
}

func NewJwtManager(expirationToken string, expirationRefresh string, secretKey string) *JwtManager {
	return &JwtManager{expirationToken: expirationToken, expirationRefresh: expirationRefresh, secretKey: secretKey}
}

func (jm *JwtManager) GenerateJwtAccessToken(payload PayloadClaims) (accessToken string, err error) {
	expirationAccessToken, _ := strconv.Atoi(jm.expirationToken)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid.New().String(),
		"type": "access",
		"exp":  time.Now().Add(time.Minute * time.Duration(expirationAccessToken)).Unix(),
		"sub":  payload,
	}).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) GenerateJwtRefreshToken(payload PayloadClaims) (refreshToken string, err error) {
	expirationRefreshToken, _ := strconv.Atoi(jm.expirationRefresh)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid.New().String(),
		"type": "refresh",
		"exp":  time.Now().Add(time.Hour * 24 * time.Duration(expirationRefreshToken)).Unix(),
		"sub":  payload,
	}).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) VerifyJwtToken(accessToken string) (claims jwt.Claims, err error) {
	parsedToken, err := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"})).
		Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
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

func (jm *JwtManager) ExtractJwtToken(headerToken string) (clearedToken string, err error) {
	splitToken := strings.Split(headerToken, "Bearer")
	if len(splitToken) != 2 {
		return "", TokenWithout
	}

	spaceToken := strings.TrimSpace(splitToken[1])
	if len(spaceToken) < 1 {
		return "", TokenWithout
	}
	return spaceToken, nil
}
