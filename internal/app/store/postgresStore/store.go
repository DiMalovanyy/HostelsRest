package postgresStore

import (
	"database/sql"

	"github.com/UniverOOP/internal/app/store"
)

type Store struct {
	db                *sql.DB
	userRepository    *UserRepo
	facultyRepository *FucultyRepo
	hostelRepository  *HostelRepo
	roomRepository    *RoomRepo
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepo {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepo{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Faculty() store.FacultyRepo {
	if s.facultyRepository != nil {
		return s.facultyRepository
	}
	s.facultyRepository = &FucultyRepo{
		store: s,
	}

	return s.facultyRepository
}

func (s *Store) Hostel() store.HostelRepo {
	if s.hostelRepository != nil {
		return s.hostelRepository
	}
	s.hostelRepository = &HostelRepo{
		store: s,
	}

	return s.hostelRepository
}

func (s *Store) Room() store.RoomRepo {
	if s.roomRepository != nil {
		return s.roomRepository
	}

	s.roomRepository = &RoomRepo{
		store: s,
	}

	return s.roomRepository
}
