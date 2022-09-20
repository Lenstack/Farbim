package infrastructure

import (
	"fmt"
	"github.com/Lenstack/farm_management/internal/core/application"
	desc "github.com/Lenstack/farm_management/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServer struct {
}

// NewGrpcServer Create a Grpc Server
func NewGrpcServer(port string, microservices application.MicroserviceServer) *GrpcServer {
	fmt.Printf("Grpc Server Is Running In Port: %s\n", port)
	listenServer, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("%s", err)
	}

	grpcServer := grpc.NewServer()

	desc.RegisterMicroserviceServer(grpcServer, &microservices)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listenServer)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &GrpcServer{}
}
