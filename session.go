package goserver

// Session reflects a session entry in the database
type Session struct {
	ID      int    `db:"id"`
	UserID  int    `db:"userid"`
	Token   string `db:"token"`
	Origin  string `db:"origin"`
	Expires string `db:"expires"` // time of expiration in string form
}
