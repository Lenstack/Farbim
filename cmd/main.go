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
		GrpcPort                 = viper.Get("GRPC_PORT").(string)
		PostgresDatabaseName     = viper.Get("POSTGRES_DATABASE_NAME").(string)
		PostgresDatabaseUser     = viper.Get("POSTGRES_DATABASE_USER").(string)
		PostgresDatabasePassword = viper.Get("POSTGRES_DATABASE_PASSWORD").(string)
		JwtExpirationToken       = viper.Get("JWT_EXPIRATION_TOKEN").(string)
		JwtExpirationRefresh     = viper.Get("JWT_EXPIRATION_REFRESH").(string)
		JwtSecret                = viper.Get("JWT_SECRET").(string)
		ApiVersion               = viper.Get("API_VERSION").(string)
		ApiPort                  = viper.Get("API_PORT").(string)
		RedisHost                = viper.Get("REDIS_HOST").(string)
		RedisPort                = viper.Get("REDIS_PORT").(string)
		RedisPassword            = viper.Get("REDIS_PASSWORD").(string)
	)

	postgres := infrastructure.NewPostgres(
		PostgresHost, PostgresPort, PostgresDatabaseName,
		PostgresDatabaseUser, PostgresDatabasePassword,
	)

	redisManager := infrastructure.NewRedisManager(RedisHost, RedisPort, RedisPassword)

	tokenManager := utils.NewJwtManager(JwtExpirationToken, JwtExpirationRefresh, JwtSecret)

	//Register Services
	redisService := services.NewRedisService(redisManager.Client)
	userService := services.NewUserService(postgres.Database)
	authenticationService := services.NewAuthenticationService(postgres.Database, *tokenManager, redisManager.Client)

	//Register Http Handlers
	middlewareApplication := application.NewMiddlewareApplication(*userService, *redisService, *tokenManager)
	authenticationApplication := application.NewAuthenticationApplication(*authenticationService, *tokenManager)
	userApplication := application.NewUserApplication(*userService)

	microservices := application.NewMicroserviceServer(
		*middlewareApplication,
		*userApplication,
		*authenticationApplication,
	)

	go func() {
		infrastructure.NewGrpcServer(GrpcPort)
	}()

	router := infrastructure.NewRouter(*microservices, ApiVersion)
	infrastructure.NewHttpServer(ApiPort, router.App)
}
