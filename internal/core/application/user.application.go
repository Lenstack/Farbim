package application

import (
	"encoding/json"
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type IUserApplication interface {
	Show(writer http.ResponseWriter, request *http.Request)
	ShowBy(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Destroy(writer http.ResponseWriter, request *http.Request)
}

type UserApplication struct {
	userService services.UserService
}

func NewUserApplication(userService services.UserService) *UserApplication {
	return &UserApplication{userService: userService}
}

func (ua *UserApplication) Show(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
	users, err := ua.userService.Show()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest,
			Message: utils.Message(err.Error()), Errors: err})
		return
	}

	if len(users) == 0 {
		writer.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(writer).Encode(utils.ResponseSuccess{Items: utils.NOT_ITEMS})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(utils.ResponseSuccess{Page: 0, PerPage: 0, TotalItems: int64(len(users)), Items: users})
}

func (ua *UserApplication) ShowBy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
	userId := chi.URLParam(request, "id")

	user, err := ua.userService.ShowBy(userId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest,
			Message: utils.Message(err.Error()), Errors: err})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(utils.ResponseSuccess{Items: user})
}

func (ua *UserApplication) Update(writer http.ResponseWriter, request *http.Request) {
	newUser := entities.User{}

	writer.Header().Add("Content-Type", "application-json")
	userId := chi.URLParam(request, "id")
	_ = json.NewDecoder(request.Body).Decode(&newUser)

	errValidate := utils.ValidateStruct(newUser)
	if errValidate != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Errors: errValidate})
		return
	}

	user, err := ua.userService.Update(userId, newUser)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest,
			Message: utils.Message(err.Error()), Errors: err})
		return
	}
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(utils.ResponseSuccess{Code: http.StatusOK, Message: utils.UPDATED, Items: user.Id})
}

func (ua *UserApplication) Destroy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
	userId := chi.URLParam(request, "id")

	rowsAffected, err := ua.userService.Destroy(userId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest,
			Message: utils.Message(err.Error()), Errors: err})
		return
	}

	if rowsAffected == 0 {
		fmt.Println(rowsAffected)
		writer.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest, Message: utils.NOT_EXIST})
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(utils.ResponseSuccess{Code: http.StatusOK, Message: utils.DELETED})
}
