package database

import "crud/model"

var users []model.User

func Add(u model.User) {
	users = append(users, u)
}

func Slice() []model.User {
	return users
}
