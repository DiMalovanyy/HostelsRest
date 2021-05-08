package postgresStore

import (
	"database/sql"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type UserRepo struct {
	store *Store
}

func (repo *UserRepo) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	return repo.store.db.QueryRow(
		"INSERT INTO users (name, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		user.Name,
		user.Email,
		user.EncryptedPassword,
	).Scan(&user.Id)
}

func (repo *UserRepo) FindByName(name string) (*model.User, error) {
	u := &model.User{}

	if err := repo.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password FROM users WHERE name = $1", name).Scan(
		&u.Id, &u.Name, &u.Email, &u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := repo.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password FROM users WHERE email = $1", email).Scan(
		&u.Id, &u.Name, &u.Email, &u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
