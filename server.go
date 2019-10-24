package goserver

import "net/http"

// Server is the handle on the main server instance
type Server struct {
	mux      *http.ServeMux
	users    UserRepoSqlite3
	sessions SessionRepoSqlite3
}

// NewServer correctly instantiates a new server
func NewServer(u UserRepoSqlite3, s SessionRepoSqlite3) *Server {
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

	// Check for authorization

	s.mux.ServeHTTP(w, req)
}

func (s *Server) getSecretStuff(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) userCreate(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) userChangePasswd(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) userAuthenticate(w http.ResponseWriter, req *http.Request) {

}
