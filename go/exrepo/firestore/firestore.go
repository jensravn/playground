package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/jensravn/playground/go/exrepo/entity"
)

type UserRepo struct {
}

const usersCollection = "Users"

var client *firestore.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	var err error
	client, err = firestore.NewClient(ctx, "jensravn-goplayground")
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}
}

func (r *UserRepo) GetUser(id int) (*entity.User, error) {
	q := client.Collection("users").Where("ID", "==", id)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		user := &entity.User{
			ID:   doc.Data()["ID"].(int),
			Name: doc.Data()["Name"].(string),
		}
		return user, nil
	}
}

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepo) CreateUser(user entity.User) error {
	_, _, err := client.Collection("users").Add(ctx, user)
	if err != nil {
		log.Fatalf("Failed adding user: %v - err: %v", user, err)
	}
	return err
}

func (r *UserRepo) UpdateUser(user entity.User) error {
	panic("not implemented") // TODO: Implement
}

func (r *UserRepo) DeleteUser(id int) error {
	panic("not implemented") // TODO: Implement
}
