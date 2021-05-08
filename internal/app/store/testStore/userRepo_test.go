package testStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
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
	assert.NotEqual(t, len(testRepo.userByEmail), 0)
	assert.NotEqual(t, len(testRepo.userByName), 0)
	assert.Equal(t, len(testRepo.userByEmail), len(testRepo.userByName))
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
