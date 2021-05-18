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
		"INSERT INTO users (name, email, encrypted_password, sex, room_id, grade, faculty_id ) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.Name,
		user.Email,
		user.EncryptedPassword,
		model.UNDEFINED,
		0, 0, 0,
	).Scan(&user.Id)
}

func (repo *UserRepo) FindByName(name string) (*model.User, error) {
	u := &model.User{}

	if err := repo.store.db.QueryRow(
		"SELECT id, name, email, encrypted_password, sex, room_id, grade, faculty_id FROM users WHERE name = $1", name).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
		&u.Sex,
		&u.RoomId,
		&u.Grade,
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
		"SELECT id, name, email, encrypted_password, sex, room_id, grade, faculty_id FROM users WHERE email = $1", email).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
		&u.Sex,
		&u.RoomId,
		&u.Grade,
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
		"SELECT id, name, email, encrypted_password, sex, room_id, grade, faculty_id FROM users WHERE id=$1", id).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.EncryptedPassword,
		&u.Sex,
		&u.RoomId,
		&u.Grade,
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

func (repo *UserRepo) GetAllUsersByRoomId(roomId int) ([]*model.User, error) {
	users := make([]*model.User, 0)

	rows, err := repo.store.db.Query(
		"SELECT id, name, email, encrypted_password, sex, room_id, grade, faculty_id FROM users WHERE room_id = $1", roomId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.EncryptedPassword,
			&user.Sex,
			&user.RoomId,
			&user.Grade,
			&user.FacultyId,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, store.ErrEmptyData
	}

	return users, nil
}
