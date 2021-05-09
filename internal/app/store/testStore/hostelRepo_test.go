package testStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestHostelRepo_CreateHostel(t *testing.T) {
	s := New()
	facRepo := s.Faculty()
	hostelRepo := s.Hostel()
	h := model.TestHostel(t)
	fac := model.TestFaculty(t)
	facRepo.CreateFaculty(fac)
	fac, _ = facRepo.GetFacultyByName(fac.Name)
	h.FacultyId = fac.Id

	assert.NoError(t, hostelRepo.CreateHostel(h))
	assert.NotEqual(t, len(hostelRepo.(*HostelRepo).Hostels), 0)
}

func TestHostelRepo_GetHostelsByfucultyId(t *testing.T) {

	// s := New()
	// facRepo := s.Faculty()
	// hostelRepo := s.Hostel()
	// h := model.TestHostel(t)
	// fac := model.TestFaculty(t)
	// facRepo.CreateFaculty(fac)
	// fac, _ = facRepo.GetFacultyByName(fac.Name)
	// h.FacultyId = fac.Id

	// _, err := hostelRepo.GetHostelsByFucultyId(fac.Id + 1)
	// assert.EqualError(t, err, store.ErrEmptyData.Error())

	// hs, err := hostelRepo.GetHostelsByFucultyId(fac.Id)
	// assert.NoError(t, err)
	// assert.NotEqual(t, len(hs), 0)
	// assert.Equal(t, hs[0].Description, h.Description)
}
