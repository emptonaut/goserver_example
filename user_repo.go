package goserver

// UserRepo is a general interface definition for getting persistent user data
type UserRepo interface {
	CreateUser(user *User) error
	UpdateUserPasswd(user *User) error
	GetUserByID(user *User) error
	GetUserByUsername(user *User) error
}
