package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) DisableAccount(_ context.Context, request *pkg.DisableAccountRequest) (*pkg.DisableAccountResponse, error) {
	err := ms.AuthenticationService.DisableAccount(request.Id)
	if err != nil {
		return nil, err
	}
	return &pkg.DisableAccountResponse{Message: "the account has been deactivated"}, nil
}
