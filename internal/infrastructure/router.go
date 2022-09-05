package infrastructure

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/go-chi/chi/v5"
)

type Router struct{}

func NewRouter(microservices application.MicroServer) *chi.Mux {
	app := chi.NewRouter()
	app.Group(func(router chi.Router) {
		router.Get("/user", microservices.List)
		router.Get("/user/{id}", microservices.ShowBy)
		router.Post("/user", microservices.Create)
		router.Put("/user/{id}", microservices.Update)
		router.Delete("/user/{id}", microservices.Destroy)
	})
	return app
}
