package utils

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"strings"
	"time"
)

type IJwtManager interface {
	GenerateJwtAccessToken(userId string) (accessToken string, err error)
	GenerateJwtRefreshToken(userId string) (refreshToken string, err error)
	VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error)
	ExtractJwtToken(headerToken string) (clearedToken string, err error)
	GetSubClaims(claims jwt.MapClaims) (userId string, err error)
}

type JwtManager struct {
	expirationAccess  string
	expirationRefresh string
	secretKey         string
}

func NewJwtManager(expirationAccess string, expirationRefresh string, secretKey string) *JwtManager {
	return &JwtManager{expirationAccess: expirationAccess, expirationRefresh: expirationRefresh, secretKey: secretKey}
}

func (jm *JwtManager) GenerateJwtAccessToken(userId string, expiration string) (accessToken string, err error) {
	expirationAccessToken, _ := strconv.Atoi(expiration)
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * time.Duration(expirationAccessToken)).Unix(),
		"sub": userId,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) GenerateJwtRefreshToken(userId string, expiration string) (refreshToken string, err error) {
	expirationRefreshToken, _ := strconv.Atoi(expiration)
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(expirationRefreshToken)).Unix(),
		"sub": userId,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jm.secretKey))
}

func (jm *JwtManager) VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"}))
	parsedToken, err := parser.Parse(accessToken, func(token *jwt.Token) (any, error) {
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

func (jm *JwtManager) GetSubClaims(claims jwt.MapClaims) (userId string, err error) {
	marshalClaims, err := json.Marshal(claims["sub"])
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(marshalClaims, &userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
