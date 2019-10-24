package goserver

import (
	"github.com/jmoiron/sqlx"
)

type SessionRepoSqlite3 struct {
	db *sqlx.DB
}

func NewSessionRepoSqlite3(db *sqlx.DB) (*SessionRepoSqlite3, error) {

	return nil, nil
}

func (repo *SessionRepoSqlite3) CreateSession(s *Session) error {

	return nil
}

func (repo *SessionRepoSqlite3) DeleteSession(s *Session) error {

	return nil
}

func (repo *SessionRepoSqlite3) GetSession(s *Session) error {

	return nil
}
