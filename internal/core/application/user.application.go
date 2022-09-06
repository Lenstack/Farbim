package application

import (
	"encoding/json"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/utils"
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
		_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest, Message: utils.Message(err.Error()), Errors: err})
	}
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(utils.ResponseSuccess{Page: 0, PerPage: 0, TotalItems: int64(len(users)), Items: users})
}

func (ua *UserApplication) ShowBy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
}

func (ua *UserApplication) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
}

func (ua *UserApplication) Destroy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application-json")
}
