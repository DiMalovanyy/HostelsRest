package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type Store struct {
	userRepository *UserRepo
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepo {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepo{
		userByName:  make(map[string]*model.User),
		userByEmail: make(map[string]*model.User),
	}

	return s.userRepository
}
