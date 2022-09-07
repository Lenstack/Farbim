package main

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/infrastructure"
	"os"
)

func main() {
	infrastructure.Load()
	sqlite := infrastructure.NewSqlite(os.Getenv("SQLITE_DATASOURCE"))
	infrastructure.NewMigration(sqlite.Database)

	//Register Services
	userService := services.NewUserService(sqlite.Database)
	authenticationService := services.NewAuthenticationService(sqlite.Database)

	//Register Http Handlers
	authenticationApplication := application.NewAuthenticationApplication(*authenticationService)
	userApplication := application.NewUserApplication(*userService)

	microservices := application.NewMicroserviceServer(
		*userApplication,
		*authenticationApplication,
	)
	router := infrastructure.NewRouter(*microservices)
	infrastructure.NewHttpServer(os.Getenv("API_PORT"), router)
}
