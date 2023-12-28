package repo

import "github.com/jensravn/playground/go/cmd/repository/entity"

type UserRepo interface {
	GetUser(id int) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
}

var Test = "Test"
