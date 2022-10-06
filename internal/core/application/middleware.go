package application

import (
	"encoding/json"
	"errors"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"strings"
)

type IMiddlewareApplication interface {
	GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	HttpInterceptor() runtime.ServeMuxOption
	HttpErrorInterceptor() runtime.ServeMuxOption
	AuthorizationHttpInterceptor(serverMux *runtime.ServeMux) http.HandlerFunc
	isAuthorized(ctx context.Context, method string, interceptorType string, request *http.Request) (err error)
	isAccessibleRoles() map[string][]string
}

type MiddlewareApplication struct {
	jwtManager utils.JwtManager
}

func NewMiddlewareApplication(jwtManager utils.JwtManager) *MiddlewareApplication {
	return &MiddlewareApplication{jwtManager: jwtManager}
}

var (
	GRPC = "GRPC"
	HTTP = "HTTP"
)

func (ma *MiddlewareApplication) GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	err := ma.isAuthorized(ctx, info.FullMethod, GRPC, nil)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (ma *MiddlewareApplication) GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	err := ma.isAuthorized(stream.Context(), info.FullMethod, GRPC, nil)
	if err != nil {
		return err
	}
	return handler(srv, stream)
}

func (ma *MiddlewareApplication) HttpInterceptor() runtime.ServeMuxOption {
	httpServerOptions := runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		log.Println("--> http interceptor: ", req.URL, req.Method)

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

	return httpServerOptions
}

func (ma *MiddlewareApplication) HttpErrorInterceptor() runtime.ServeMuxOption {
	httpErrorServerOptions := runtime.WithErrorHandler(
		func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			log.Println("--> http-error interceptor: ", request.URL, request.Method)

			writer.Header().Add("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			_ = marshaler.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest, Errors: err.Error()})
		})
	return httpErrorServerOptions
}

func (ma *MiddlewareApplication) AuthorizationHttpInterceptor(serverMux *runtime.ServeMux) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("--> authorization interceptor: ", request.URL, request.Method)

		err := ma.isAuthorized(nil, request.RequestURI, HTTP, request)
		if err != nil {
			writer.Header().Add("Content-Type", "application-json")
			writer.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusBadRequest, Errors: err.Error()})
			return
		}
		serverMux.ServeHTTP(writer, request)
	}
}

func (ma *MiddlewareApplication) isAuthorized(ctx context.Context, method string, interceptorType string, request *http.Request) (err error) {
	accessibleRoles, ok := ma.isAccessibleRoles()[method]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "this method not been registered for access")
	}
	// can access without roles
	if len(accessibleRoles) == 0 {
		return nil
	}

	var headerToken, accessToken string

	if interceptorType == GRPC {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		values := md.Get("authorization")
		if len(values) == 0 {
			return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		headerToken = strings.Join(values, "")
		accessToken, err = ma.jwtManager.ExtractJwtToken(headerToken)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, err.Error())
		}
	}

	if interceptorType == HTTP {
		accessToken, err = ma.jwtManager.ExtractJwtToken(request.Header.Get("Authorization"))
		if err != nil {
			return err
		}
	}

	claims, err := ma.jwtManager.VerifyJwtToken(accessToken)
	if err != nil {
		return err
	}

	subClaims := claims["sub"].(map[string]interface{})
	listRoles := subClaims["Roles"].([]interface{})

	for _, accessibleRole := range accessibleRoles {
		for _, roles := range listRoles {
			cleanListRoles := strings.Split(roles.(string), ",")
			for _, role := range cleanListRoles {
				if accessibleRole == role {
					return nil
				}
			}
		}
	}

	if interceptorType == HTTP {
		return errors.New("no permission to access this Method")
	}

	return status.Errorf(codes.Unauthenticated, "no permission to access this RPC")
}

func (ma *MiddlewareApplication) isAccessibleRoles() map[string][]string {
	const httpPath = "/v1/"
	const servicePath = "/microservice.Microservice/"
	const serviceReflectionPath = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"

	return map[string][]string{
		serviceReflectionPath:               {},
		servicePath + "SignUp":              {"Admin"},
		servicePath + "SignIn":              {},
		servicePath + "Logout":              {"Admin"},
		servicePath + "VerifyAccount":       {"Admin"},
		servicePath + "DisableAccount":      {"Admin"},
		servicePath + "GetUsers":            {},
		servicePath + "GetUser":             {"Admin"},
		servicePath + "UpdateUserPassword":  {"Admin"},
		servicePath + "DeleteUser":          {"Admin"},
		servicePath + "UpdateProfile":       {"Admin"},
		servicePath + "CreateFarm":          {"Admin"},
		servicePath + "CreateCategory":      {"Admin"},
		servicePath + "GetCategories":       {"Admin"},
		servicePath + "GetCategory":         {"Admin"},
		httpPath + "user":                   {},
		httpPath + "authentication/sign_up": {},
	}
}
