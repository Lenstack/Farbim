package repositories

import (
	"github.com/Lenstack/farm_management/internal/core/entities"
	"github.com/Masterminds/squirrel"
)

type ISessionRepository interface {
	CreateSession(session entities.Session) (sessionId string, err error)
}

type SessionRepository struct {
	Database squirrel.StatementBuilderType
}

func (sr *SessionRepository) CreateSession(session entities.Session) (sessionId string, err error) {
	bq := sr.Database.
		Insert(entities.SessionTableName).
		Columns("Id", "Type", "Token", "ClientOrigin", "ExpiredAt", "UserId").
		Values(session.Id, session.Type, session.Token, session.ClientOrigin, session.ExpiredAt, session.UserId).
		Suffix("RETURNING Id")

	err = bq.QueryRow().Scan(&sessionId)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}
