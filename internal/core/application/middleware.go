package application

import (
	"fmt"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type MiddlewareApplication struct {
	jwtManager utils.JwtManager
}

type validatable interface {
	Validate() error
}

func NewMiddlewareApplication(jwtManager utils.JwtManager) *MiddlewareApplication {
	return &MiddlewareApplication{jwtManager: jwtManager}
}

func (m *MiddlewareApplication) GrpcUnaryInterceptor() grpc.ServerOption {
	grpcUnaryServerOptions := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if v, ok := req.(validatable); ok {
			err := v.Validate()
			if err != nil {
				return nil, err
			}
		}

		err = m.isAuthorized(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	})
	return grpcUnaryServerOptions
}

func (m *MiddlewareApplication) GrpcStreamInterceptor() grpc.ServerOption {
	grpcStreamServerOptions := grpc.StreamInterceptor(func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		err = m.isAuthorized(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, stream)
	})
	return grpcStreamServerOptions
}

func (m *MiddlewareApplication) HttpInterceptor() runtime.ServeMuxOption {
	httpServerOptions := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		return nil
	})
	return httpServerOptions
}

func (m *MiddlewareApplication) isAuthorized(ctx context.Context, method string) (err error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	headerToken := strings.Join(values, "")
	accessToken, err := m.jwtManager.ExtractJwtToken(headerToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	claims, err := m.jwtManager.VerifyJwtToken(accessToken)
	if err != nil {
		return err
	}
	fmt.Println(claims)
	return status.Errorf(codes.Unauthenticated, "no permission to access this RPC")
}
