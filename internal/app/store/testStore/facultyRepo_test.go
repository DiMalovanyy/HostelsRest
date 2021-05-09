package testStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestFucultyRepo_Create(t *testing.T) {
	s := New()
	repo := s.Faculty()
	f := model.TestFaculty(t)

	assert.NoError(t, repo.CreateFaculty(f))

	facultyRepo := repo.(*FacultyRepo)
	assert.NotEqual(t, len(facultyRepo.Faculties), 0)
}

func TestFucultyRepo_GetAllFaculties(t *testing.T) {
	s := New()
	repo := s.Faculty()
	f := model.TestFaculty(t)

	_, err := repo.GetAllFaculties()
	assert.EqualError(t, err, store.ErrEmptyData.Error())

	repo.CreateFaculty(f)

	fs, err := repo.GetAllFaculties()
	assert.NoError(t, err)
	assert.NotEqual(t, len(fs), 0)
	assert.Equal(t, fs[0].Name, f.Name)

}
