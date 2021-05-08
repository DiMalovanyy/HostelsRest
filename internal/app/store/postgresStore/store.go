package postgresStore

import (
	"database/sql"

	"github.com/UniverOOP/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepo
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
