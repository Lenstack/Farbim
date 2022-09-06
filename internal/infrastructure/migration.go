package infrastructure

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"gorm.io/gorm"
)

type Migration struct {
	database *gorm.DB
}

func NewMigration(database *gorm.DB) *Migration {
	err := database.AutoMigrate(&entities.User{})
	if err != nil {
		return nil
	}
	return &Migration{database: database}
}
