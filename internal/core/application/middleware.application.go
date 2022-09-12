package application

import (
	"encoding/json"
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/utils"
	"net/http"
)

type IMiddlewareApplication interface {
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
	headerToken := request.Header.Get("Authorization")

	extractToken, err := m.jwtManager.ExtractJwtToken(headerToken)
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

func (m *MiddlewareApplication) ProtectedRoutes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application-json")
		headerToken := request.Header.Get("Authorization")

		extractToken, err := m.jwtManager.ExtractJwtToken(headerToken)
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
