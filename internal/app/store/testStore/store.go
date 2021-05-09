package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type Store struct {
	userRepository    *UserRepo
	facultyRepository *FacultyRepo
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepo {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepo{
		userById: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Faculty() store.FacultyRepo {
	if s.facultyRepository != nil {
		return s.facultyRepository
	}

	s.facultyRepository = &FacultyRepo{
		Faculties: make(map[int]*model.Faculty),
	}

	return s.facultyRepository
}
