package application

import (
	"encoding/json"
	"github.com/Lenstack/farm_management/internal/core/services"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type IMiddlewareApplication interface {
	GrpcAuthorization(ctx context.Context, method string) error
	HttpAuthorization(request *http.Request, method string) error
}

type MiddlewareApplication struct {
	jwtManager        utils.JwtManager
	loggerManager     *zap.Logger
	roleService       services.RoleService
	permissionService services.PermissionService
}

func NewMiddlewareApplication(jwtManager utils.JwtManager, loggerManager *zap.Logger, roleService services.RoleService, permissionService services.PermissionService) *MiddlewareApplication {
	return &MiddlewareApplication{jwtManager: jwtManager, loggerManager: loggerManager, roleService: roleService, permissionService: permissionService}
}

func (ma *MiddlewareApplication) GrpcAuthorization(ctx context.Context, method string) error {
	permissions, ok := ma.Permissions()[method]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "this method not been registered for access")
	}
	if len(permissions) == 0 {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	headerToken := strings.Join(values, "")
	accessToken, err := ma.jwtManager.ExtractJwtToken(headerToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	claims, err := ma.jwtManager.VerifyJwtToken(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	userIdByClaims, err := ma.jwtManager.GetSubClaims(claims)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	err = ma.VerifyPermissions(userIdByClaims, permissions)
	if err != nil {
		return err
	}
	return nil
}

func (ma *MiddlewareApplication) HttpAuthorization(serverMux *runtime.ServeMux) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ma.loggerManager.Sugar().Infof("-> HttpAuthorization %s <-", request.RequestURI)
		err := ma.HasPermission(request)
		if err != nil {
			writer.Header().Add("Content-Type", "application-json")
			writer.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(writer).Encode(utils.ResponseError{Code: http.StatusUnauthorized, Errors: err.Error()})
			return
		}
		serverMux.ServeHTTP(writer, request)
	}
}

func (ma *MiddlewareApplication) HasPermission(request *http.Request) error {
	permissions, ok := ma.Permissions()[request.RequestURI]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "this method not been registered for access")
	}
	if len(permissions) == 0 {
		return nil
	}

	accessToken, err := ma.jwtManager.ExtractJwtToken(request.Header.Get("Authorization"))
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	claims, err := ma.jwtManager.VerifyJwtToken(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	userIdByClaims, err := ma.jwtManager.GetSubClaims(claims)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	err = ma.VerifyPermissions(userIdByClaims, permissions)
	if err != nil {
		return err
	}
	return nil
}

func (ma *MiddlewareApplication) Permissions() map[string][]string {
	listPermissions := make(map[string][]string)
	permissions, _ := ma.permissionService.GetPermissions()

	for _, permission := range permissions {
		listPermissions[permission.Service] = permission.Roles
	}
	return listPermissions
}

func (ma *MiddlewareApplication) VerifyPermissions(userId string, permissions []string) error {
	roles, _ := ma.roleService.GetUserRolesId(userId)
	if len(roles) == 0 {
		return status.Errorf(codes.Unauthenticated, "you have no assigned roles")
	}
	for _, permission := range permissions {
		for _, role := range roles {
			if permission == role {
				return nil
			}
		}
	}
	return status.Errorf(codes.Unauthenticated, "you don't have the right role to access this RPC")
}
