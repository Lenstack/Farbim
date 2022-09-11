package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Router struct {
	App        *chi.Mux
	JwtManager utils.JwtManager
}

var jwtUtil utils.JwtManager

func NewRouter(microservices application.MicroserviceServer, apiVersion string, jwtManager utils.JwtManager) *Router {
	jwtUtil = jwtManager

	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.AllowContentType("application/json"))
	app.Use(middleware.CleanPath)

	version := fmt.Sprintf("/%s", apiVersion)

	app.Route(version+"/user", func(userRouter chi.Router) {
		userRouter.Use(ProtectedRoutesMiddleware)
		userRouter.Get("/", microservices.UserApplication.Show)
		userRouter.Get("/{id}", microservices.UserApplication.ShowBy)
		userRouter.Put("/{id}", microservices.UserApplication.Update)
		userRouter.Delete("/{id}", microservices.UserApplication.Destroy)
	})

	app.Route(version+"/authentication", func(authenticationRouter chi.Router) {
		authenticationRouter.Post("/sign_in", microservices.AuthenticationApplication.SignIn)
		authenticationRouter.Post("/sign_up", microservices.AuthenticationApplication.SignUp)
		authenticationRouter.Post("/logout", microservices.AuthenticationApplication.Logout)
	})

	return &Router{App: app}
}

func ProtectedRoutesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		extractToken, err := jwtUtil.ExtractJwtToken(request)
		if err != nil {
			writer.Header().Add("Content-Type", "application-json")
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		_, err = jwtUtil.VerifyJwtToken(extractToken)
		if err != nil {
			writer.Header().Add("Content-Type", "application-json")
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		next.ServeHTTP(writer, request)
	})
}
