package services

import (
	"fmt"

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

func (db *usersDB) GetByKey(key, value string) (schema.User, error) {
	k := fmt.Sprintf("%s:%s", key, value)
	user, ok := db.users[k]
	if !ok {
		return schema.User{}, app.ErrNotFound
	}

	return user, nil
}

func (db *usersDB) GetAll() ([]schema.User, error) {
	users := make([]schema.User, len(db.users))
	return users, nil
}

func (db *usersDB) Insert(email, fullname string) (schema.User, error) {
	emailKey := fmt.Sprintf("email:%s", email)
	//fullnameKey := fmt.Sprintf("fullname:%s", fullname)
	u := schema.User{
		Email:    email,
		Fullname: fullname,
	}
	db.users[emailKey] = u
	fmt.Println("len after insert = ", len(db.users))
	//db.users[fullnameKey] = &u
	return u, nil
}
