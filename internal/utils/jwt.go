package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IJwtManager interface {
	GenerateJwtToken(payload interface{}) (token string, err error)
	VerifyJwtToken(accessToken string) (claims jwt.MapClaims, err error)
	RefreshJwtToken(payload interface{}) (token string, err error)
	ExtractJwtToken(request *http.Request) (err error)
}

type JwtManager struct {
	expirationToken   string
	expirationRefresh string
	secretKey         string
}

func NewJwtManager(expirationToken string, expirationRefresh string, secretKey string) *JwtManager {
	return &JwtManager{expirationToken: expirationToken, expirationRefresh: expirationRefresh, secretKey: secretKey}
}

func (jm *JwtManager) GenerateJwtToken(payload interface{}) (token string, err error) {
	expiration, _ := strconv.Atoi(jm.expirationToken)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid.New().String(),
		"exp":  time.Now().Add(time.Minute * time.Duration(expiration)).Unix(),
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

func (jm *JwtManager) ExtractJwtToken(request *http.Request) (clearedToken string, err error) {
	headerToken := request.Header.Get("Authorization")
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

func (jm *JwtManager) RefreshJwtToken(payload interface{}) (token string, err error) {
	expiration, _ := strconv.Atoi(jm.expirationRefresh)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": uuid.New().String(),
		"exp":  time.Now().Add(time.Hour * 24 * time.Duration(expiration)).Unix(),
		"sub":  payload,
	}).SignedString([]byte(jm.secretKey))
}
