package app

import "github.com/gtarcea/1DevDayTalk2014/schema"

type UsersService interface {
	CreateUser(email, fullname string) (schema.User, error)
	GetAll() ([]schema.User, error)
	GetUserByEmail(email string) (schema.User, error)
}

type UserDBService interface {
	GetByEmail(email string) (schema.User, error)
	GetAll() ([]schema.User, error)
	Insert(email, fullname string) (schema.User, error)
}
