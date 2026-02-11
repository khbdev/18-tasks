package create

import (
	"crud/database"
	"crud/model"
)

var lastid int = 0

func CreateDatabase(name, email string) {
	lastid++
	user := model.User{
		Id:    lastid,
		Name:  name,
		Email: email,
	}
	database.Add(user)

}
