package application

import (
	"encoding/json"
	"github.com/Lenstack/farm_management/internal/core/services"
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
	writer.Header().Set("Content-Type", "application-json")
	users, err := ua.userService.Show()
	if err != nil {
		_ = json.NewEncoder(writer).Encode(&err)
		writer.WriteHeader(http.StatusBadRequest)
	}
	_ = json.NewEncoder(writer).Encode(&users)
	writer.WriteHeader(http.StatusOK)
}

func (ua *UserApplication) ShowBy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")
}

func (ua *UserApplication) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")
}

func (ua *UserApplication) Destroy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application-json")
}
