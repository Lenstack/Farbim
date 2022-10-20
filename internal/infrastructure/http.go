package infrastructure

import (
	"github.com/Lenstack/farm_management/internal/core/application"
	"github.com/Lenstack/farm_management/internal/utils"
	desc "github.com/Lenstack/farm_management/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
)

type HttpServer struct {
	port string
}

func NewHttpServer(apiVersion string, port string, loggerManager *zap.Logger, microservices application.MicroserviceServer) {
	serverMux := runtime.NewServeMux(HttpServerOptions()...)
	err := desc.RegisterMicroserviceHandlerServer(context.Background(), serverMux, &microservices)
	if err != nil {
		loggerManager.Sugar().Errorf("%s", err)
	}

	muxRouter := NewRouter(microservices, apiVersion, loggerManager, serverMux)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		loggerManager.Sugar().Errorf("%s", err)
	}

	loggerManager.Sugar().Infof("starting HTTP server in %s", listener.Addr().String())
	if err := http.Serve(listener, muxRouter); err != nil {
		loggerManager.Sugar().Errorf("%s", err)
	}
}

func HttpServerOptions() (opts []runtime.ServeMuxOption) {
	opts = append(opts, HttpMetadataInterceptor())
	opts = append(opts, HttpJsonOptionInterceptor())
	opts = append(opts, HttpErrorInterceptor())
	return opts
}

func HttpJsonOptionInterceptor() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
}

func HttpMetadataInterceptor() runtime.ServeMuxOption {
	metadataOptions := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		md := make(map[string]string)
		md["authorization"] = req.Header.Get("Authorization")
		if method, ok := runtime.RPCMethod(ctx); ok {
			md["method"] = method
		}
		if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
			md["pattern"] = pattern
		}
		return metadata.New(md)
	})

	return metadataOptions
}

func HttpErrorInterceptor() runtime.ServeMuxOption {
	return runtime.WithErrorHandler(
		func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			_ = marshaler.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest, Errors: err.Error()})
		})
}
