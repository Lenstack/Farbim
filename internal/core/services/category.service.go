package services

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/Masterminds/squirrel"
)

type ICategoryService interface {
	CreateCategory(category entities.Category) (categoryId string, err error)
	GetCategories() (categories []entities.Category, err error)
	GetCategory(categoryId string) (category entities.Category, err error)
}

type CategoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(database squirrel.StatementBuilderType) *CategoryService {
	return &CategoryService{
		categoryRepository: repositories.CategoryRepository{
			Database: database,
		},
	}
}

func (cs *CategoryService) CreateCategory(category entities.Category) (categoryId string, err error) {
	return cs.categoryRepository.CreateCategory(category)
}

func (cs *CategoryService) GetCategories() (categories []entities.Category, err error) {
	return cs.categoryRepository.GetCategories()
}

func (cs *CategoryService) GetCategory(categoryId string) (category entities.Category, err error) {
	return cs.categoryRepository.GetCategory(categoryId)
}
