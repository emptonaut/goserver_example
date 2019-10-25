
SERVER_SRC := server.go $(wildcard session*) $(wildcard user*) server/main.go
SERVER_OUT := bin/server
DBNAME := dummy.db

$(SERVER_OUT): $(SERVER_SRC)
	mkdir -p bin
	go build -o $(SERVER_OUT) ./server

.PHONY: db
db:
	rm -rf dummy.db
	sqlite3 $(DBNAME) << db_setup.sql
