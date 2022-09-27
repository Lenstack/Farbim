package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) DeleteUser(_ context.Context, request *pkg.DeleteUserRequest) (*pkg.DeleteUserResponse, error) {
	_, err := ms.UserService.Destroy(request.Id)
	if err != nil {
		return nil, err
	}
	return &pkg.DeleteUserResponse{Message: "this item has been deleted successfully"}, nil
}
