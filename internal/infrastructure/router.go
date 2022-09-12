package infrastructure

import (
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	App *chi.Mux
}

func NewRouter(microservices application.MicroserviceServer, apiVersion string) *Router {
	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.AllowContentType("application/json"))
	app.Use(middleware.CleanPath)

	version := fmt.Sprintf("/%s", apiVersion)

	//Protected Routes
	app.Group(func(appRouter chi.Router) {
		appRouter.Use(microservices.MiddlewareApplication.ProtectedRoutes)
		appRouter.Route(version+"/user", func(userRouter chi.Router) {
			userRouter.Get("/", microservices.UserApplication.Show)
			userRouter.Get("/{id}", microservices.UserApplication.ShowBy)
			userRouter.Put("/{id}", microservices.UserApplication.Update)
			userRouter.Delete("/{id}", microservices.UserApplication.Destroy)
		})

		appRouter.Route(version+"/authorization", func(authorizationRouter chi.Router) {
			authorizationRouter.Post("/refresh_token", microservices.MiddlewareApplication.RefreshToken)
		})
	})

	app.Route(version+"/authentication", func(authenticationRouter chi.Router) {
		authenticationRouter.Post("/sign_in", microservices.AuthenticationApplication.SignIn)
		authenticationRouter.Post("/sign_up", microservices.AuthenticationApplication.SignUp)
		authenticationRouter.Post("/logout", microservices.AuthenticationApplication.Logout)
	})

	return &Router{App: app}
}
