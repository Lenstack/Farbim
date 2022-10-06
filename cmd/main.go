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
	userService := services.NewUserService(postgres.Database)
	authenticationService := services.NewAuthenticationService(postgres.Database, *tokenManager, redisManager.Client)
	profileService := services.NewProfileService(postgres.Database)
	farmService := services.NewFarmService(postgres.Database)
	categoryService := services.NewCategoryService(postgres.Database)

	middlewareApplication := application.NewMiddlewareApplication(*tokenManager)

	microservices := application.NewMicroserviceServer(
		*middlewareApplication,
		*authenticationService,
		*userService,
		*profileService,
		*farmService,
		*categoryService,
	)

	go func() {
		infrastructure.NewGrpcServer(GrpcPort, *microservices)
	}()

	infrastructure.NewHttpServer(ApiVersion, ApiPort, *microservices)
}
