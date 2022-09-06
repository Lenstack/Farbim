package application

import (
	"encoding/json"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/utils"
	"net/http"
)

type IAuthenticationApplication interface {
	SignIn(writer http.ResponseWriter, request *http.Request)
	SignUp(writer http.ResponseWriter, request *http.Request)
	Logout(writer http.ResponseWriter, request *http.Request)
}

type AuthenticationApplication struct {
	authenticationService services.AuthenticationService
}

func NewAuthenticationApplication(authenticationService services.AuthenticationService) *AuthenticationApplication {
	return &AuthenticationApplication{authenticationService: authenticationService}
}

func (aa *AuthenticationApplication) SignIn(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")

	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(
		utils.ResponseSuccess{Code: http.StatusCreated, Message: utils.SIGNIN},
	)
}

func (aa *AuthenticationApplication) SignUp(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")
	user := entities.User{}
	_ = json.NewDecoder(request.Body).Decode(&user)

	errValidate := utils.ValidateStruct(&user)
	if errValidate != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Errors: errValidate})
		return
	}

	err := aa.authenticationService.SignUp(user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(
			utils.ResponseError{Errors: err},
		)
	}

	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(
		utils.ResponseSuccess{Code: http.StatusCreated, Message: utils.SIGNUP},
	)
}

func (aa *AuthenticationApplication) Logout(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")

	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(
		utils.ResponseSuccess{Code: http.StatusCreated, Message: utils.LOGOUT},
	)
}
