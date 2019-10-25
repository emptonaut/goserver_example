package goserver

type UserRepo interface {
	CreateUser(user *User) error
	UpdateUserPasswd(user *User) error
	GetUser(user *User) error
}
