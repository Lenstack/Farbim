package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) VerifyEmail(ctx context.Context, request *pkg.VerifyEmailRequest) (*pkg.VerifyEmailResponse, error) {
	_, err := ms.AuthenticationService.VerifyEmail(request.Token)
	if err != nil {
		return nil, err
	}
	return &pkg.VerifyEmailResponse{Message: "this email has been verified successfully"}, nil
}
