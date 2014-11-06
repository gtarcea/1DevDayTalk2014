package services

import (
	"github.com/gtarcea/1DevDayTalk2014/app"
	"github.com/gtarcea/1DevDayTalk2014/schema"
)

type usersDB struct {
	users map[string]schema.User
}

func NewUsersDB() app.UserDBService {
	return &usersDB{
		users: make(map[string]schema.User),
	}
}

func (db *usersDB) GetByEmail(email string) (schema.User, error) {
	user, ok := db.users[email]
	if !ok {
		return schema.User{}, app.ErrNotFound
	}

	return user, nil
}

func (db *usersDB) GetAll() ([]schema.User, error) {
	users := make([]schema.User, 0, len(db.users))
	i := 0
	for _, user := range db.users {
		users = append(users, user)
		i++
	}
	return users, nil
}

func (db *usersDB) Insert(email, fullname string) (schema.User, error) {
	u := schema.User{
		Email:    email,
		Fullname: fullname,
	}
	db.users[email] = u
	return u, nil
}
