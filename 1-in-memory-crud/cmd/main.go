package main

import (
	"crud/create"
	"crud/getall"
	"crud/getid"
	"fmt"
)

func main() {

	create.CreateDatabase("Azizbek", "Xasanov@gmail.com")

	users := getall.GetAll()

	fmt.Println(users)

	user, ok := getid.GetId(1)

	if !ok {
		fmt.Println("user topilmadi")
		return
	}

	fmt.Println(user)

}
