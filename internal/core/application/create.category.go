package application

import (
	"errors"
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) CreateCategory(_ context.Context, request *pkg.CreateCategoryRequest) (*pkg.CreateCategoryResponse, error) {
	newCategory := entities.Category{Name: request.GetName()}
	if request.GetName() == "" {
		return nil, errors.New("category name is empty")
	}
	_, err := ms.CategoryService.CreateCategory(newCategory)
	if err != nil {
		return nil, err
	}
	return &pkg.CreateCategoryResponse{Message: "category has been created successfully."}, nil
}
