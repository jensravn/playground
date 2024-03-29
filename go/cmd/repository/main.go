package main

import (
	"log"

	"github.com/jensravn/playground/go/cmd/repository/entity"
	"github.com/jensravn/playground/go/cmd/repository/firestore"
)

func main() {
	println("Hello REPOSITORY playground!")

	// userRepo := mockdb.UserRepo{}
	userRepo := firestore.UserRepo{}

	userRepo.CreateUser(entity.User{ID: 1, Name: "John"})

	user, err := userRepo.GetUser(1)
	if err != nil {
		log.Fatalf("GetUser(): %+v", err)
	}
	log.Printf("Hello: %+v", user)

	user2, err := userRepo.GetUser(2)
	if err != nil {
		log.Fatalf("GetUser(): %+v", err)
	}
	log.Printf("Hello: %+v", user2)
}
