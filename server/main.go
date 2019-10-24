package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"

	shoeserv "github.com/shoelick/goserver_example"
)

func main() {

	// init DB

	// setup db repositories

	// setup server and go
	s := shoeserv.NewServer()
	http.HandleFunc("/", s)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
