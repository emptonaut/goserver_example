package goserver

// SessionRepo defines the basic interface for managing authentication sessions
type SessionRepo interface {
	Create(s *Session) error
	DeleteByToken(s *Session) error
	GetByID(s *Session) error
	GetByToken(s *Session) error
}
