package goserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Server is the handle on the main server instance
type Server struct {
	mux      *http.ServeMux
	users    UserRepo
	sessions SessionRepo
}

// NewServer correctly instantiates a new server
func NewServer(u UserRepo, s SessionRepo) *Server {
	server := &Server{
		mux:      http.NewServeMux(),
		users:    u,
		sessions: s,
	}
	server.Init()
	return server
}

// Init sets up http endpoints
func (s *Server) Init() {

	s.mux.HandleFunc("/secrets", s.getSecretStuff)

	s.mux.HandleFunc("/user/create", s.userCreate)
	s.mux.HandleFunc("/user/ChangePasswd", s.userCreate)
	s.mux.HandleFunc("/user/Authenticate", s.userCreate)
}

// Handle enables our server to receive requests and route them as we see fit
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Get the request body
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnavailableForLegalReasons)
		return
	}

	// Check for authorization
	data := &RequestData{}
	if err := json.Unmarshal(reqBody, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.mux.ServeHTTP(w, req)
}

func (s *Server) getSecretStuff(w http.ResponseWriter, req *http.Request) {
	secrets := "42"
	w.Write([]byte(secrets))
}

func (s *Server) userCreate(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) userChangePasswd(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) userAuthenticate(w http.ResponseWriter, req *http.Request) {

}
