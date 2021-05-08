package store

import "github.com/UniverOOP/internal/app/model"

type UserRepo interface {
	CreateUser(*model.User) error
	FindByName(string) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
