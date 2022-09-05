package databases

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Sqlite struct {
	datasource string
	Database   *gorm.DB
}

func NewSqlite(datasource string) *Sqlite {
	database, err := gorm.Open(sqlite.Open(datasource), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &Sqlite{Database: database}
}
