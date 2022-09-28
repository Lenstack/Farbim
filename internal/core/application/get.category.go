package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) GetCategory(ctx context.Context, request *pkg.GetCategoryRequest) (*pkg.GetCategoryResponse, error) {
	return &pkg.GetCategoryResponse{Id: ""}, nil
}
