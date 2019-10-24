package goserver

type SessionRepo interface {
	CreateSession(s *Session) error
	DeleteSession(s *Session) error
	GetSession(s *Session) error
}
