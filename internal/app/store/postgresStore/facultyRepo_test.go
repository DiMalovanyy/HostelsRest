package postgresStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestFacultyRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("faculty")

	s := New(db)

	f := model.TestFaculty(t)
	assert.NoError(t, s.Faculty().CreateFaculty(f))
	assert.NotNil(t, f)
}

func TestFacultyRepository_GetAllFaculties(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("faculty")

	s := New(db)
	f := model.TestFaculty(t)

	_, err := s.Faculty().GetAllFaculties()
	assert.EqualError(t, err, store.ErrEmptyData.Error())

	s.Faculty().CreateFaculty(f)
	fs, err := s.Faculty().GetAllFaculties()
	assert.NoError(t, err)
	assert.NotEqual(t, len(fs), 0)
	assert.Equal(t, fs[0].Name, f.Name)

}

func TestFucultyRepo_GetFucultyByName(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("faculty")

	s := New(db)
	f := model.TestFaculty(t)
	repo := s.Faculty()

	_, err := repo.GetFacultyByName(f.Name)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	repo.CreateFaculty(f)

	fac, err := repo.GetFacultyByName(f.Name)
	assert.NoError(t, err)
	assert.NotNil(t, fac)
	assert.Equal(t, fac.Name, f.Name)
}
