package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) VerifyAccount(_ context.Context, request *pkg.VerifyAccountRequest) (*pkg.VerifyAccountResponse, error) {
	err := ms.AuthenticationService.VerifyAccount(request.Id)
	if err != nil {
		return nil, err
	}
	return &pkg.VerifyAccountResponse{Message: "the account has been activated"}, nil
}
