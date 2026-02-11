package getall

import (
	"crud/database"
	"crud/model"
)

func GetAll() []model.User {
	user := database.Slice()
	return user
}
