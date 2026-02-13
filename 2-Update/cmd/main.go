package main

import (
	"fmt"

)



type User struct {
	ID int 
	Name string
}

var (
	users []User
	lastId = 0	
)

func CreateUser(name string) {
	lastId++
	user := User {
		ID: lastId,
		Name: name,
	}
	users = append(users, user)
}


func UpdateUser(id int, name string) {
	for i, u := range users {
		if u.ID == id {
			users[i].Name = name
			return
		}
	}
}


func main(){
CreateUser("Azizbek")
fmt.Println(users)
UpdateUser(1, "Javohir")
fmt.Println(users)
}