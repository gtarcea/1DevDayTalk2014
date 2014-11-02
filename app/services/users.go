package services

import (
	"github.com/gtarcea/1DevDayTalk2014/app"
	"github.com/gtarcea/1DevDayTalk2014/schema"
)

type usersService struct {
	db app.UserDBService
}

func NewUsers(db app.UserDBService) app.UsersService {
	return &usersService{
		db: db,
	}
}

func (u *usersService) CreateUser(email, fullname string) (schema.User, error) {
	return u.db.Insert(email, fullname)
}

func (u *usersService) GetUserByEmail(email string) (schema.User, error) {
	return u.db.GetByKey("email", email)
}
