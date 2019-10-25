package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
	shoeserv "github.com/shoelick/goserver_example"
)

func main() {

	// init DB
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	//_, err = db.Exec("read db_setup.sql")
	if err != nil {
		log.Fatal(err)
	}

	// setup db repositories
	var userRepo *shoeserv.UserRepoSqlite3
	userRepo, err = shoeserv.NewUserRepoSqlite3(db)
	if err != nil {
		log.Fatal(err)
	}
	var sessionRepo *shoeserv.SessionRepoSqlite3
	sessionRepo, err = shoeserv.NewSessionRepoSqlite3(db)
	if err != nil {
		log.Fatal(err)
	}

	// setup server and go
	s := shoeserv.NewServer(userRepo, sessionRepo)
	http.Handle("/", s)
	err = http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
