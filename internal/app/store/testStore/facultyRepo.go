package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type FacultyRepo struct {
	Faculties map[int]*model.Faculty
}

func (f *FacultyRepo) CreateFaculty(faculty *model.Faculty) error {

	//TODO: Some Validation (OR ???)

	faculty.Id = len(f.Faculties) + 1
	f.Faculties[faculty.Id] = faculty

	return nil

}

func (f *FacultyRepo) GetAllFaculties() ([]*model.Faculty, error) {

	if len(f.Faculties) == 0 {
		return nil, store.ErrEmptyData
	}

	faculties := make([]*model.Faculty, 0)
	for _, fac := range f.Faculties {
		faculties = append(faculties, fac)
	}

	return faculties, nil
}
