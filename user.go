package goserver

// User reflects a user entry in the database
// This implementation doesn't actually use the salt field because of the bcrypt
// package provided by Go.
type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Salt     string `db:"salt"`
}
