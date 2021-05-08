package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type UserRepo struct {
	userByName  map[string]*model.User //Name - user
	userByEmail map[string]*model.User //Name - email
}

func (repo *UserRepo) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}
	repo.userByName[user.Name] = user
	repo.userByEmail[user.Email] = user
	user.Id = len(repo.userByName)
	return nil
}

func (repo *UserRepo) FindByName(name string) (*model.User, error) {
	u, ok := repo.userByName[name]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	u, ok := repo.userByEmail[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil

}
