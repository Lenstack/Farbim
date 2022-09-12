package main

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/infrastructure"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/spf13/viper"
)

func main() {
	infrastructure.Load()

	var (
		PostgresHost             = viper.Get("POSTGRES_HOST").(string)
		PostgresPort             = viper.Get("POSTGRES_PORT").(string)
		PostgresDatabaseName     = viper.Get("POSTGRES_DATABASE_NAME").(string)
		PostgresDatabaseUser     = viper.Get("POSTGRES_DATABASE_USER").(string)
		PostgresDatabasePassword = viper.Get("POSTGRES_DATABASE_PASSWORD").(string)
		JwtExpirationToken       = viper.Get("JWT_EXPIRATION_TOKEN").(string)
		JwtExpirationRefresh     = viper.Get("JWT_EXPIRATION_REFRESH").(string)
		JwtSecret                = viper.Get("JWT_SECRET").(string)
		ApiVersion               = viper.Get("API_VERSION").(string)
		ApiPort                  = viper.Get("API_PORT").(string)
	)

	postgres := infrastructure.NewPostgres(
		PostgresHost, PostgresPort, PostgresDatabaseName,
		PostgresDatabaseUser, PostgresDatabasePassword,
	)

	tokenManager := utils.NewJwtManager(JwtExpirationToken, JwtExpirationRefresh, JwtSecret)

	//Register Services
	userService := services.NewUserService(postgres.Database)
	authenticationService := services.NewAuthenticationService(postgres.Database, *tokenManager)

	//Register Http Handlers
	middlewareApplication := application.NewMiddlewareApplication(*userService, *tokenManager)
	authenticationApplication := application.NewAuthenticationApplication(*authenticationService, *tokenManager)
	userApplication := application.NewUserApplication(*userService)

	microservices := application.NewMicroserviceServer(
		*middlewareApplication,
		*userApplication,
		*authenticationApplication,
	)
	router := infrastructure.NewRouter(*microservices, ApiVersion)
	infrastructure.NewHttpServer(ApiPort, router.App)
}
