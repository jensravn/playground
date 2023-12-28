package mockdb

import (
	"fmt"

	"github.com/jensravn/playground/go/cmd/repository/entity"
)

type UserRepo struct {
	users []entity.User
}

func (r *UserRepo) GetUser(id int) (*entity.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	return r.users, nil
}

func (r *UserRepo) CreateUser(user entity.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *UserRepo) UpdateUser(user entity.User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}
	return nil
}

func (r *UserRepo) DeleteUser(id int) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return nil
}
