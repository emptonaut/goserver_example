package goserver

import "github.com/jmoiron/sqlx"

type UserRepoSqlite3 struct {
	db *sqlx.DB
}

func NewUserRepoSqlite3(db *sqlx.DB) (*UserRepoSqlite3, error) {

	return nil, nil
}

func (repo *UserRepoSqlite3) CreateUser(s *User) error {

	return nil
}

func (repo *UserRepoSqlite3) UpdateUserPasswd(s *User) error {

	return nil
}

func (repo *UserRepoSqlite3) GetUser(s *User) error {

	return nil
}
