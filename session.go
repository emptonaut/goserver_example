package goserver

type Session struct {
	ID      int
	UserID  int
	Token   string
	Origin  string
	Expires string
}
