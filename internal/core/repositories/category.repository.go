package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type ICategoryRepository interface {
	CreateCategory(category entities.Category) (categoryId string, err error)
	GetCategories() (categories []entities.Category, err error)
	GetCategory(categoryId string) (category entities.Category, err error)
}

type CategoryRepository struct {
	Database squirrel.StatementBuilderType
}

func (cr *CategoryRepository) CreateCategory(category entities.Category) (categoryId string, err error) {
	category.Id = uuid.New().String()
	bq := cr.Database.
		Insert(entities.CategoryTableName).
		Columns("Id", "Name").
		Values(category.Id, category.Name).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&categoryId)
	if err != nil {
		return "", utils.ErrorManager(err)
	}

	return categoryId, nil
}

func (cr *CategoryRepository) GetCategories() (categories []entities.Category, err error) {
	category := entities.Category{}
	bq := cr.Database.
		Select("id", "name", "createdAt", "updatedAt").
		From(entities.CategoryTableName)

	rows, err := bq.Query()
	if err != nil {
		return nil, utils.ErrorManager(err)
	}

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.Name,
			&category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return nil, utils.ItemWithout
	}

	return categories, nil
}

func (cr *CategoryRepository) GetCategory(categoryId string) (category entities.Category, err error) {
	bq := cr.Database.Select("Id", "Name", "CreatedAt", "UpdatedAt").
		From(entities.CategoryTableName).
		Where(squirrel.Eq{"id": categoryId})

	err = bq.QueryRow().Scan(&category.Id, &category.Name,
		&category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return entities.Category{}, utils.ErrorManager(err)
	}
	return category, nil
}
