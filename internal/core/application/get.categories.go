package application

import (
	"github.com/Lenstack/farm_management/pkg"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (ms *MicroserviceServer) GetCategories(_ context.Context, _ *emptypb.Empty) (*pkg.GetCategoriesResponse, error) {
	categories, err := ms.CategoryService.GetCategories()
	if err != nil {
		return nil, err
	}

	categoriesResponse := make([]*pkg.GetCategoryResponse, 0, len(categories))
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, &pkg.GetCategoryResponse{
			Id:        category.Id,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreatedAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		})
	}

	return &pkg.GetCategoriesResponse{Page: 1, PerPage: 10, Total: int32(len(categoriesResponse)),
		Categories: categoriesResponse}, nil
}
