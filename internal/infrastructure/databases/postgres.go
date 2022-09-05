package databases

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	host         string
	port         string
	databaseName string
	user         string
	password     string
	Database     *gorm.DB
}

func NewPostgres(host string, port string, databaseName string, user string, password string) *Postgres {
	datasource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		host, user, password, databaseName, port)
	database, err := gorm.Open(postgres.Open(datasource), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &Postgres{Database: database}
}
