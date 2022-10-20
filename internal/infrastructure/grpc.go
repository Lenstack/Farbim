package infrastructure

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/pkg"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type GrpcServer struct {
}

func NewGrpcServer(port string, loggerManager *zap.Logger, microservices application.MicroserviceServer) *GrpcServer {
	listenServer, err := net.Listen("tcp", ":"+port)
	if err != nil {
		loggerManager.Sugar().Errorf("%s", err)
	}
	loggerManager.Sugar().Infof("starting GRPC server in %s", listenServer.Addr().String())

	grpcServer := grpc.NewServer(GrpcServerOptions(loggerManager, microservices)...)
	pkg.RegisterMicroserviceServer(grpcServer, &microservices)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listenServer)
	if err != nil {
		loggerManager.Sugar().Errorf("%s", err)
	}
	return &GrpcServer{}
}

func GrpcServerOptions(loggerManager *zap.Logger, microservices application.MicroserviceServer) (opts []grpc.ServerOption) {
	loggerOptions := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	opts = append(opts, grpc.ChainUnaryInterceptor(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(loggerManager, loggerOptions...),
		GrpcUnaryInterceptor(loggerManager, microservices),
	))
	opts = append(opts, grpc.ChainStreamInterceptor(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_zap.StreamServerInterceptor(loggerManager, loggerOptions...),
		GrpcStreamInterceptor(loggerManager, microservices),
	))
	return opts
}

func GrpcUnaryInterceptor(loggerManager *zap.Logger, microservices application.MicroserviceServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		loggerManager.Sugar().Infof("-> GrpcUnaryInterceptor %s <-", info.FullMethod)
		err = microservices.MiddlewareApplication.GrpcAuthorization(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func GrpcStreamInterceptor(loggerManager *zap.Logger, microservices application.MicroserviceServer) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		loggerManager.Sugar().Infof("-> GrpcStreamInterceptor %s <-", info.FullMethod)
		err = microservices.MiddlewareApplication.GrpcAuthorization(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}
