
COMMON_SRC := common.go $(wildcard session*) $(wildcard user*)
SERVER_SRC := server.go $(COMMON_SRC) $(wildcard server/*.go)
CLIENT_SRC := client.go $(COMMON_SRC) $(wildcard client/*.go)
SERVER_OUT := bin/server
CLIENT_OUT := bin/client
DBNAME := dummy.db

.PHONY: db test

all: server client

server: $(SERVER_OUT)

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

clean:
	@rm -v $(SERVER_OUT) $(CLIENT_OUT) server.crt server.csr
