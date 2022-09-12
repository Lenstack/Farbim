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
	writer.Header().Add("Content-Type", "application-json")
	extractToken, err := m.ExtractToken(request)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}

	userId, err := m.jwtManager.VerifyJwtToken(extractToken)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}

	user, err := m.UserService.ShowBy(userId)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}

	refreshToken, err := m.jwtManager.GenerateJwtRefreshToken(user.Id)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}

	_, err = m.UserService.UpdateUserRefreshToken(user.Id, refreshToken)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(
		utils.ResponseSuccess{Code: http.StatusOK, TokenRefresh: refreshToken},
	)
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
		writer.Header().Add("Content-Type", "application-json")
		extractToken, err := m.ExtractToken(request)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		userId, err := m.jwtManager.VerifyJwtToken(extractToken)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		fmt.Println("Token UserId: ", userId)
		next.ServeHTTP(writer, request)
	})
}
