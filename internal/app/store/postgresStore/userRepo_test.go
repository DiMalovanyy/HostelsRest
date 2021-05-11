package postgresStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)

	u := model.TestUser(t)
	assert.NoError(t, s.User().CreateUser(u))
	assert.NotNil(t, u)
}

func TestUserRepo_FindByName(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	name := "some name"
	_, err := s.User().FindByName(name)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Name = name
	s.User().CreateUser(u)
	u, err = s.User().FindByName(name)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().CreateUser(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_Upgrade(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("users")

	s := New(db)

	u := model.TestUser(t)

	repo := s.User()
	repo.CreateUser(u)

	sex := model.MEN
	roomId := 1
	facultyId := 1
	grade := 1

	assert.EqualError(t, repo.Upgrade(0, sex, roomId, facultyId, grade), store.ErrRecordNotFound.Error())
	assert.NoError(t, repo.Upgrade(u.Id, sex, roomId, facultyId, grade))

	u, _ = repo.FindById(u.Id)

	assert.Equal(t, u.Sex, sex)
	assert.Equal(t, u.RoomId, roomId)
	assert.Equal(t, u.FacultyId, facultyId)

}
