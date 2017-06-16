package usecases

import (
	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
)

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}

type User struct {
	Id       int
	IsAdmin  bool
	Customer entities.Customer
}
