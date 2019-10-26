package goserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type endpoint struct {
	handler      func(http.ResponseWriter, *http.Request)
	authRequired bool
}

// Server is the handle on the main server instance
type Server struct {
	mux       *http.ServeMux
	users     UserRepo
	sessions  SessionRepo
	endpoints map[string]*endpoint
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

	s.endpoints = make(map[string]*endpoint)

	s.endpoints["/user/create"] = &endpoint{s.userCreate, false}
	s.endpoints["/user/authenticate"] = &endpoint{s.userAuthenticate, false}
	s.endpoints["/secrets"] = &endpoint{s.getSecretStuff, true}
	s.endpoints["/user/changePasswd"] = &endpoint{s.userChangePasswd, true}

	for key, val := range s.endpoints {
		s.mux.HandleFunc(key, val.handler)
	}
}

// Handle enables our server to receive requests and route them as we see fit
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Get the request body
	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnavailableForLegalReasons)
		return
	}

	if _, ok := s.endpoints[req.URL.Path]; ok {

		// Check for authorization
		//data := &RequestData{}
		//if err := json.Unmarshal(reqBody, data); err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		// First see if we need to have authentication for this endpoint

		// If yes, check for a token

		// If no token, return not authorized

		// Otherwise, lookup the token and see if it's valid

		// If valid, allow through

		s.mux.ServeHTTP(w, req)
	} else {
		http.NotFoundHandler().ServeHTTP(w, req)
	}
}

func (s *Server) getSecretStuff(w http.ResponseWriter, req *http.Request) {
	secrets := "42"
	w.Write([]byte(secrets))
}

func (s *Server) userCreate(w http.ResponseWriter, req *http.Request) {
	data, err := parseRequestData(w, req)
	if err != nil {
		return
	}

	if data.Username == "" || data.Password == "" {
		http.Error(w, "must provide username and password", http.StatusExpectationFailed)
		return
	}

	hash, err := hashPasswd([]byte(data.Password))
	if err != nil {
		http.Error(w, "password hash failed", http.StatusExpectationFailed)
		return
	}

	// Create user
	user := &User{
		Username: data.Username,
		Password: hash,
	}
	s.users.CreateUser(user)

}

func (s *Server) userChangePasswd(w http.ResponseWriter, req *http.Request) {
	_, err := parseRequestData(w, req)
	if err != nil {
		return
	}

}

func (s *Server) userAuthenticate(w http.ResponseWriter, req *http.Request) {
	_, err := parseRequestData(w, req)
	if err != nil {
		return
	}

}

// The better solution to something like this would be to rewrite the http.Handler interface to accept
// a third argument for common data that every endpoint would require, such as session data.
// However, that's beyond the scope of this example exercise. For now, every endpoint must
// reparse the RequestData even if the root receive (ServerHTTP) already parsed it. Known design flaw.
func parseRequestData(w http.ResponseWriter, req *http.Request) (*RequestData, error) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return nil, err
	}
	data := &RequestData{}
	err = json.Unmarshal(reqBody, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return nil, err
	}
	return data, err
}

func hashPasswd(passwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(passwd, bcrypt.MinCost) // TODO check cost?
	return string(hash), err
}
