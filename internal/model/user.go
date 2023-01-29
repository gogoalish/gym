package model

type User struct {
	Id       string
	Email    string
	Name     string
	Password string
}

type UserRepository interface {
	CreateUser(*User) error
}
