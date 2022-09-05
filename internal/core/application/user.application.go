package application

import (
	"encoding/json"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type IUserApplication interface {
	List(writer http.ResponseWriter, request *http.Request)
	ShowBy(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Destroy(writer http.ResponseWriter, request *http.Request)
	Verify(writer http.ResponseWriter, request *http.Request)
}

type UserApplication struct {
	UserService services.UserService
}

func NewUserApplication(userService services.UserService) *UserApplication {
	return &UserApplication{UserService: userService}
}

func (ua *UserApplication) List(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	users, err := ua.UserService.List()
	if err != nil {
		return
	}
	_ = json.NewEncoder(writer).Encode(&users)
}

func (ua *UserApplication) ShowBy(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(request, "id")
	userId, _ := strconv.ParseInt(id, 10, 64)
	user, err := ua.UserService.ShowBy(userId)
	if err != nil {
		return
	}
	_ = json.NewEncoder(writer).Encode(user)
}

func (ua *UserApplication) Create(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	user := entities.User{}
	_ = json.NewDecoder(request.Body).Decode(&user)
	err := ua.UserService.Create(user)
	if err != nil {
		return
	}
	_ = json.NewEncoder(writer).Encode(user)
}

func (ua *UserApplication) Update(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ua *UserApplication) Destroy(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ua *UserApplication) Verify(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
