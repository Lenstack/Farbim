package main

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/infrastructure"
	"github.com/Lenstack/farm_management/internal/infrastructure/databases"
	"os"
)

func main() {
	infrastructure.Load()

	sqlite := databases.NewSqlite(os.Getenv("SQLITE_DATASOURCE"))
	/*
		postgres := databases.NewPostgres(
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DATABASE_NAME"), os.Getenv("POSTGRES_DATABASE_USER"),
			os.Getenv("POSTGRES_DATABASE_PASSWORD"))
	*/
	databases.NewMigration(sqlite.Database)
	//databases.NewMigration(postgres.Database)

	userService := services.NewUserService(sqlite.Database)
	userApplication := application.NewUserApplication(*userService)
	microservices := application.NewMicroServer(*userApplication)

	router := infrastructure.NewRouter(*microservices)
	infrastructure.NewHttpServer(os.Getenv("API_PORT"), router)
}
