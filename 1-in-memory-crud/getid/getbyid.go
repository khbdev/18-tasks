package getid

import (
	"crud/getall"
	"crud/model"
)

func GetId(id int) (model.User, bool) {
	users := getall.GetAll()

	for _, u := range users {
		if u.Id == id {
			return u, true
		}
	}
	return model.User{}, false
}
