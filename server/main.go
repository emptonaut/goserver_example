package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
	shoe "github.com/shoelick/goserver_example"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)

	// init DB
	db, err := sqlx.Open("sqlite3", "dummy.db")
	if err != nil {
		log.Fatal(err)
	}
	//_, err = db.Exec("read db_setup.sql")
	if err != nil {
		log.Fatal(err)
	}

	// setup db repositories
	var userRepo *shoe.UserRepoSqlite3
	userRepo, err = shoe.NewUserRepoSqlite3(db)
	if err != nil {
		log.Fatal(err)
	}
	var sessionRepo *shoe.SessionRepoSqlite3
	sessionRepo, err = shoe.NewSessionRepoSqlite3(db)
	if err != nil {
		log.Fatal(err)
	}

	// setup server and go
	s := shoe.NewServer(userRepo, sessionRepo)
	http.Handle("/", s)
	err = http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
