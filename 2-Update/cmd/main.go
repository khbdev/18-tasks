package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000000000000000000; i++ {
			fmt.Println("A:", i)
		}
	}()

	go func() {
		defer wg.Done()
		for j := 0; j < 10000000000000; j++ {
			fmt.Println("B:", j)
		}
	}()

	wg.Wait()
}