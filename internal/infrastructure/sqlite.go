package infrastructure

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Sqlite struct {
	datasource string
	Database   squirrel.StatementBuilderType
}

func NewSqlite(datasource string) *Sqlite {
	database, err := sql.Open("sqlite3", datasource)
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &Sqlite{Database: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(database)}
}
