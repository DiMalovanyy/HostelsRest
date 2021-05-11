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
		"INSERT INTO users (name, email, encrypted_password, sex, room_id, faculty_id ) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		user.Name,
		user.Email,
		user.EncryptedPassword,
		model.UNDEFINED,
		0, 0,
	).Scan(&user.Id)
}

func (repo *UserRepo) FindByName(name string) (*model.User, error) {
	u := &model.User{}

	if err := repo.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password, sex, room_id, faculty_id FROM users WHERE name = $1", name).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
		&u.Sex,
		&u.RoomId,
		&u.FacultyId,
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
		"SELECT id, name, email, encrypted_password, sex, room_id, faculty_id FROM users WHERE email = $1", email).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
		&u.Sex,
		&u.RoomId,
		&u.FacultyId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (repo *UserRepo) FindById(id int) (*model.User, error) {
	u := &model.User{}

	if err := repo.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password, sex, room_id, faculty_id FROM users WHERE id=$1", id).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
		&u.Sex,
		&u.RoomId,
		&u.FacultyId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (repo *UserRepo) Upgrade(userId int, sex model.Sex, roomId int, facultyId int, grade int) error {

	u, err := repo.FindById(userId)
	if err != nil {
		return err
	}

	if err := repo.store.db.QueryRow(
		"UPDATE users SET sex = $1, room_id = $2, faculty_id = $3, grade = $4 WHERE id = $5 RETURNING sex, room_id, faculty_id, grade",
		sex, roomId, facultyId, grade, userId).Scan(
		&u.Sex,
		&u.RoomId,
		&u.FacultyId,
		&u.Grade,
	); err != nil {
		return err
	}

	return nil
}
