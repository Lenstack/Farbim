package infrastructure

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"moul.io/chizap"
)

func NewRouter(microservices application.MicroserviceServer, apiVersion string, loggerManager *zap.Logger, mux *runtime.ServeMux) *chi.Mux {
	app := chi.NewRouter()
	app.Use(chizap.New(loggerManager, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))
	app.Use(middleware.AllowContentType("application/json"))
	app.Use(middleware.CleanPath)
	app.Use(cors.Handler(cors.Options{AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"}}))
	app.HandleFunc("/"+apiVersion+"/*", microservices.MiddlewareApplication.HttpAuthorization(mux))
	return app
}
