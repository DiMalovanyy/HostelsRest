package testStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	s := New()
	u := model.TestUser(t)

	repo := s.User()
	assert.NoError(t, repo.CreateUser(u))
	assert.NotNil(t, u)

	testRepo, ok := repo.(*UserRepo)
	assert.True(t, ok)
	assert.NotEqual(t, len(testRepo.userById), 0)
	assert.NotEqual(t, u.Id, 0)
}

func TestUserRepo_FindByName(t *testing.T) {
	s := New()

	email := "example@gmail.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().CreateUser(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByEmail(t *testing.T) {
	s := New()

	name := "some name"
	_, err := s.User().FindByName(name)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Name = name
	s.User().CreateUser(u)
	u, err = s.User().FindByName(name)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepo_FindById(t *testing.T) {
	s := New()

	id := 1
	_, err := s.User().FindById(id)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Id = id
	s.User().CreateUser(u)
	u, err = s.User().FindById(id)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepo_Upgrade(t *testing.T) {
	u := model.TestUser(t)
	s := New()

	repo := s.User()
	repo.CreateUser(u)

	sex := model.MEN
	roomId := 1
	facultyId := 1

	assert.EqualError(t, repo.Upgrade(0, sex, roomId, facultyId), store.ErrRecordNotFound.Error())
	assert.NoError(t, repo.Upgrade(u.Id, sex, roomId, facultyId))

	assert.Equal(t, u.Sex, sex)
	assert.Equal(t, u.RoomId, roomId)
	assert.Equal(t, u.FacultyId, facultyId)
}
