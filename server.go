package goserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
	s.endpoints["/secret"] = &endpoint{s.getSecretStuff, true}
	s.endpoints["/user/changePasswd"] = &endpoint{s.userChangePasswd, true}
	s.endpoints["/user/logout"] = &endpoint{s.userLogout, true}

	for key, val := range s.endpoints {
		s.mux.HandleFunc(key, val.handler)
	}
}

// Handle enables our server to receive requests and route them as we see fit
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if endpoint, ok := s.endpoints[req.URL.Path]; ok {

		// First see if we need to have authentication for this endpoint
		if endpoint.authRequired {
			// Get the request body
			reqBody, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Error(err)
				http.Error(w, err.Error(), http.StatusUnavailableForLegalReasons)
				return
			}

			// We parse the request body here and at the beginning of every
			// endpoint.
			// The better solution to something like this would be to rewrite
			// the http.Handler interface to accept. A third argument for common
			// data that every endpoint would require, such as session data.
			// However, that's beyond the scope of this example exercise. For
			// now, every endpoint must
			// reparse the RequestData even if the root receive (ServerHTTP)
			// already parsed it. Known design flaw.
			data := &RequestData{}
			if err := json.Unmarshal(reqBody, data); err != nil {
				log.Error(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

			// Otherwise, lookup the token and see if it's valid
			authorized := true
			if data.Token == "" {
				authorized = false
			}

			session := &Session{Token: data.Token}
			if err = s.sessions.GetByToken(session); err != nil {
				log.Infof("Attempt auth with nonexistent token: %s", data.Token)
				authorized = false
			}

			if !authorized {
				http.Error(w, "please authenticate", http.StatusUnauthorized)
				return
			}
		}

		s.mux.ServeHTTP(w, req)
	} else {
		http.NotFoundHandler().ServeHTTP(w, req)
	}
}

func (s *Server) getSecretStuff(w http.ResponseWriter, req *http.Request) {
	secrets := "42"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(secrets))
}

func (s *Server) userCreate(w http.ResponseWriter, req *http.Request) {
	data, err := ParseRequestData(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	if data.Username == "" || data.Password == "" {
		http.Error(w, "must provide username and password", http.StatusExpectationFailed)
		return
	}

	hash, err := hashPasswd([]byte(data.Password))
	if err != nil {
		log.Error(err)
		http.Error(w, "password hash failed", http.StatusExpectationFailed)
		return
	}

	// Create user
	user := &User{
		Username: data.Username,
		Password: hash,
	}
	// We're assuming an enforcement of unique usernames at the database level.
	// We could check for the username before we hash, but it's probably the
	// same number of cycles or more to do a DB query
	err = s.users.CreateUser(user)
	if err != nil {
		log.Error(err)
		data.Error = err.Error()
	} else {
		log.Printf("Created user %s", user.Username)
	}

	data.Password = ""
	writeRequestData(data, w, http.StatusOK)
}

func (s *Server) userChangePasswd(w http.ResponseWriter, req *http.Request) {
	data, err := ParseRequestData(req.Body)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	// Hash new password
	hash, err := hashPasswd([]byte(data.Password))
	if err != nil {
		log.Error(err)
		http.Error(w, "password hash failed", http.StatusExpectationFailed)
		return
	}

	// Get user ID from session
	sesh := &Session{Token: data.Token}
	err = s.sessions.GetByToken(sesh)
	if err != nil {
		log.Error(err)
		return
	}

	// Require all of username and new password
	u := &User{
		ID:       sesh.UserID,
		Username: data.Username,
		Password: hash,
	}
	log.Infof("Changing password for %v", u)

	var status = http.StatusOK
	data.Password = ""
	data.Token = ""

	if err = s.users.UpdateUserPasswd(u); err != nil {
		log.Error(err)
		data.Error = err.Error()
		status = http.StatusInternalServerError
	}

	writeRequestData(data, w, status)
}

func (s *Server) userAuthenticate(w http.ResponseWriter, req *http.Request) {
	data, err := ParseRequestData(req.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	u := &User{Username: data.Username}
	if err := s.users.GetUserByUsername(u); err != nil {
		http.Error(w, "authentication failed", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.Password)); err != nil {
		http.Error(w, "authentication failed", http.StatusUnauthorized)
		return
	}

	// Password was good; create a session and send back the token
	token := uuid.New().String()
	// TODO record location for session and expire time
	session := &Session{
		Token:  token,
		UserID: u.ID,
	}

	if err := s.sessions.Create(session); err != nil {
		// Server failure
		fmt.Println(err)
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}
	log.Printf("Session created for %s", u.Username)

	// Send back the token
	data.Token = token
	data.Password = ""
	if err = writeRequestData(data, w, http.StatusOK); err != nil {
		log.Error(err)
	}
}

func (s *Server) userLogout(w http.ResponseWriter, req *http.Request) {
	data, err := ParseRequestData(req.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	sesh := &Session{Token: data.Token}
	if err := s.sessions.DeleteByToken(sesh); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
	}
	log.Infof("Logged out %s", sesh.Token)
}

func hashPasswd(passwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(passwd, bcrypt.MinCost) // TODO check cost?
	return string(hash), err
}

func writeRequestData(data *RequestData, w http.ResponseWriter, code int) error {
	bits, err := json.Marshal(data)
	if err == nil {
		w.WriteHeader(code)
		w.Write(bits)
	}
	return err
}
