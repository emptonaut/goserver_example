package goserver

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// SessionRepoSqlite3 fulfills SessionRepo using a Sqlite3 database
type SessionRepoSqlite3 struct {
	db             *sqlx.DB
	insertStmt     *sqlx.NamedStmt
	getByIDStmt    *sqlx.NamedStmt
	getByTokenStmt *sqlx.NamedStmt
	deleteStmt     *sqlx.NamedStmt
}

const (
	sessionInsert = `
		INSERT INTO sessions (userid, token, origin, expires)
		VALUES (:userid, :token, :origin, :expires)
	`

	sessionGetBase = `SELECT * FROM sessions`

	sessionGetByID    = sessionGetBase + ` WHERE id=:id LIMIT 1`
	sessionGetByToken = sessionGetBase + ` WHERE token=:token LIMIT 1`

	sessionDelete = `
		DELETE FROM sessions WHERE id=:id
	`
)

// NewSessionRepoSqlite3 prepares a repo given a sqlite db handle
func NewSessionRepoSqlite3(db *sqlx.DB) (*SessionRepoSqlite3, error) {
	var err error
	repo := &SessionRepoSqlite3{db: db}
	repo.insertStmt, err = db.PrepareNamed(sessionInsert)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", sessionInsert, err)
	}
	repo.deleteStmt, err = db.PrepareNamed(sessionDelete)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", sessionDelete, err)
	}
	repo.getByIDStmt, err = db.PrepareNamed(sessionGetByID)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", sessionGetByID, err)
	}
	repo.getByTokenStmt, err = db.PrepareNamed(sessionGetByToken)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", sessionGetByToken, err)
	}

	return repo, err
}

// Create tries to insert a session and fills in the ID of the provided Session struct
func (repo *SessionRepoSqlite3) Create(s *Session) error {
	res, err := repo.insertStmt.Exec(s)
	if err == nil {
		var id int64
		id, err = res.LastInsertId()
		if err == nil {
			s.ID = int(id)
		}
	}
	return err
}

// Delete deletes the given session using the ID field
func (repo *SessionRepoSqlite3) Delete(s *Session) (err error) {
	_, err = repo.deleteStmt.Exec(s)
	return
}

// GetByID fills the passed Session struct's fields using the ID field, which
// must be filled in
func (repo *SessionRepoSqlite3) GetByID(s *Session) (err error) {
	err = repo.getByIDStmt.Get(s, s)
	return
}

// GetByID fills the passed Session struct's fields using the ID field, which
// must be filled in
func (repo *SessionRepoSqlite3) GetByToken(s *Session) (err error) {
	err = repo.getByTokenStmt.Get(s, s)
	return
}
