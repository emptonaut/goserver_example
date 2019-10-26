
COMMON_SRC := common.go $(wildcard session*) $(wildcard user*)
SERVER_SRC := server.go server/main.go $(COMMON_SRC)
CLIENT_SRC := client.go $(COMMON_SRC)
SERVER_OUT := bin/server
CLIENT_OUT := bin/client
DBNAME := dummy.db

.PHONY: db test

server: $(SERVER_OUT)

all: server client

$(SERVER_OUT): $(SERVER_SRC)
	mkdir -p bin
	go build -o $(SERVER_OUT) ./server

client: $(CLIENT_OUT)

$(CLIENT_OUT): $(CLIENT_SRC)
	mkdir -p bin
	go build -o $(CLIENT_OUT) ./client

db:
	rm -rf dummy.db
	sqlite3 $(DBNAME) < db_setup.sql

test: db
	go test . -cover
