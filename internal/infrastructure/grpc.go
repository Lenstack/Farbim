package infrastructure

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	grpcServer *grpc.Server
}

func NewGrpcServer(port string) *GrpcServer {
	fmt.Printf("Grpc Server Is Running In Port: %s\n", port)
	listenServer, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("%s", err)
	}

	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(listenServer)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return &GrpcServer{grpcServer: grpcServer}
}
