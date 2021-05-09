package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type Store struct {
	userRepository    *UserRepo
	facultyRepository *FacultyRepo
	hostelRepository  *HostelRepo
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

func (s *Store) Hostel() store.HostelRepo {
	if s.hostelRepository != nil {
		return s.hostelRepository
	}

	s.hostelRepository = &HostelRepo{
		Hostels: make(map[int]*model.Hostel),
	}

	return s.hostelRepository
}
