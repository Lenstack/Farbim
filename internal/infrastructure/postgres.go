package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	host         string
	port         string
	databaseName string
	user         string
	password     string
	Database     squirrel.StatementBuilderType
}

func NewPostgres(host string, port string, databaseName string, user string, password string) *Postgres {
	datasource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		host, user, password, databaseName, port)
	database, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &Postgres{Database: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(database)}
}
