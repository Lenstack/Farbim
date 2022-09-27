package infrastructure

import (
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/application"
	desc "github.com/Lenstack/farm_management/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type HttpServer struct {
	port     string
	handlers *chi.Mux
}

func NewHttpServer(port string, microservices application.MicroserviceServer) {
	fmt.Printf("Http Server Is Running In Port: %s\n", port)

	serverMux := runtime.NewServeMux(microservices.MiddlewareApplication.HttpInterceptor())
	err := desc.RegisterMicroserviceHandlerServer(context.Background(), serverMux, &microservices)
	if err != nil {
		return
	}

	if err := http.ListenAndServe(":"+port, serverMux); err != nil {
		log.Fatalf("%s", err)
	}
}
