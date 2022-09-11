package application

import (
	"encoding/json"
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/utils"
	"net/http"
	"strings"
)

type IMiddlewareApplication interface {
	ExtractToken(request *http.Request) (clearedToken string, err error)
	RefreshToken(writer http.ResponseWriter, request *http.Request)
	ProtectedRoutes(next http.Handler) http.Handler
}

type MiddlewareApplication struct {
	UserService services.UserService
	jwtManager  utils.JwtManager
}

func NewMiddlewareApplication(userService services.UserService, jwtManager utils.JwtManager) *MiddlewareApplication {
	return &MiddlewareApplication{UserService: userService, jwtManager: jwtManager}
}

func (m *MiddlewareApplication) RefreshToken(writer http.ResponseWriter, request *http.Request) {
	extractToken, err := m.ExtractToken(request)
	if err != nil {
		writer.Header().Add("Content-Type", "application-json")
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}
	claims, err := m.jwtManager.VerifyJwtToken(extractToken)
	if err != nil {
		writer.Header().Add("Content-Type", "application-json")
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}

	userId := claims["sub"].(string)
	fmt.Println(userId)

	user, err := m.UserService.ShowBy(userId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

}

func (m *MiddlewareApplication) ExtractToken(request *http.Request) (clearedToken string, err error) {
	headerToken := request.Header.Get("Authorization")
	splitToken := strings.Split(headerToken, "Bearer")

	if len(splitToken) != 2 {
		return "", utils.TokenWithout
	}

	spaceToken := strings.TrimSpace(splitToken[1])
	if len(spaceToken) < 1 {
		return "", utils.TokenWithout
	}
	return spaceToken, nil
}

func (m *MiddlewareApplication) ProtectedRoutes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		extractToken, err := m.ExtractToken(request)
		if err != nil {
			writer.Header().Add("Content-Type", "application-json")
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		claims, err := m.jwtManager.VerifyJwtToken(extractToken)
		if err != nil {
			writer.Header().Add("Content-Type", "application-json")
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		fmt.Println("Token Subject: ", claims["sub"])
		next.ServeHTTP(writer, request)
	})
}
