package main

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/infrastructure"
	"github.com/Lenstack/farm_management/internal/utils"
	"os"
)

func main() {
	infrastructure.Load()
	postgres := infrastructure.NewPostgres(
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE_NAME"), os.Getenv("POSTGRES_DATABASE_USER"),
		os.Getenv("POSTGRES_DATABASE_PASSWORD"))

	tokenManager := utils.NewJwtManager(os.Getenv("JWT_EXPIRATION"), os.Getenv("JWT_SECRET"))

	//Register Services
	userService := services.NewUserService(postgres.Database)
	authenticationService := services.NewAuthenticationService(postgres.Database, *tokenManager)

	//Register Http Handlers
	authenticationApplication := application.NewAuthenticationApplication(*authenticationService)
	userApplication := application.NewUserApplication(*userService)

	microservices := application.NewMicroserviceServer(
		*userApplication,
		*authenticationApplication,
	)
	router := infrastructure.NewRouter(*microservices, os.Getenv("API_VERSION"))
	infrastructure.NewHttpServer(os.Getenv("API_PORT"), router)
}
