package goserver

// User reflects a user entry in the database
type User struct {
	ID       int
	Username string
	Password string
	Salt     string
}
