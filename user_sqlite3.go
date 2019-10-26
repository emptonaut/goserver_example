package goserver

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// UserRepoSqlite3 fulfills UserRepo using a Sqlite3 database
type UserRepoSqlite3 struct {
	db               *sqlx.DB
	insertStmt       *sqlx.NamedStmt
	updatePasswdStmt *sqlx.NamedStmt
	getByIDStmt      *sqlx.NamedStmt
}

const (
	userInsert = `
		INSERT INTO users (username, password, salt)
		VALUES (:username, :password, :salt)
	`

	userUpdatePasswd = `
		UPDATE users SET
			password=:password,
			salt=:salt
		WHERE id=:id
	`

	userGetByID = `
		SELECT * FROM users WHERE id=:id LIMIT 1
	`
)

// NewUserRepoSqlite3 prepares a repo given a sqlite db handle
func NewUserRepoSqlite3(db *sqlx.DB) (*UserRepoSqlite3, error) {

	var err error
	repo := &UserRepoSqlite3{db: db}
	repo.insertStmt, err = db.PrepareNamed(userInsert)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", userInsert, err)
	}
	repo.updatePasswdStmt, err = db.PrepareNamed(userUpdatePasswd)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", userUpdatePasswd, err)
	}
	repo.getByIDStmt, err = db.PrepareNamed(userGetByID)
	if err != nil {
		return repo, fmt.Errorf("Failed to prepare statement `%s`: %v", userGetByID, err)
	}

	return repo, err
}

// CreateUser tries to insert a user and fills in the ID of the provided user
func (repo *UserRepoSqlite3) CreateUser(s *User) error {

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

// UpdateUserPasswd attempts to update the given user's corresponding password
// and salt
func (repo *UserRepoSqlite3) UpdateUserPasswd(s *User) (err error) {
	_, err = repo.updatePasswdStmt.Exec(s)
	return
}

// GetUserByID fills the passed User struct's fields using the ID field, which
// must be filled in
func (repo *UserRepoSqlite3) GetUserByID(s *User) (err error) {
	err = repo.getByIDStmt.Get(s, s)
	return
}
