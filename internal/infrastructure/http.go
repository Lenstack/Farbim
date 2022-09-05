package infrastructure

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HttpServer struct {
	port     string
	handlers *chi.Mux
}

func NewHttpServer(port string, handlers *chi.Mux) {
	fmt.Printf("Server Is Running In Port: %s\n", port)
	if err := http.ListenAndServe(":"+port, handlers); err != nil {
		log.Fatalf("%s", err)
	}
}
