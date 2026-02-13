package main

import (
	"context"
	"fmt"
	"log"
	"unique/config"
	"unique/repostroy"

	_ "github.com/lib/pq"
)



func main(){
	db, err := config.ConfigPotgress(config.Postgress{
		Host: "localhost",
		Port: 5432,
		User: "aziz",
		Password: "aziz123",
		DBName: "azizdb",
		SSLMode: "disable",
	})
	if err != nil {
		log.Fatal("db connect err", err)
	}
	defer db.Close()

	repo := repostroy.NewUserRepostry(db)

	ctx := context.Background()

	u, err := repo.Create(ctx, "Azizbek", "Azizbek@gmail.com")
	if err != nil {
		if err == repostroy.ErrEmailAlreadyExists {
			log.Println("Email band!")
			return
		}
		log.Fatal("Create error:", err)
	}
	fmt.Println("Created", u)



}