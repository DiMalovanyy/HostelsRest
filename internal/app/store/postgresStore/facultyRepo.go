package postgresStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type FucultyRepo struct {
	store *Store
}

func (f *FucultyRepo) CreateFaculty(faculty *model.Faculty) error {
	//TODO: Some validation (OR ???)
	return f.store.db.QueryRow("INSERT INTO faculty (name) VALUES ($1) RETURNING id", faculty.Name).Scan(&faculty.Id)
}

func (f *FucultyRepo) GetAllFaculties() ([]*model.Faculty, error) {
	faculties := make([]*model.Faculty, 0)

	rows, err := f.store.db.Query("SELECT id, name FROM faculty")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		faculty := &model.Faculty{}
		if err := rows.Scan(&faculty.Id, &faculty.Name); err != nil {
			return nil, err
		}
		faculties = append(faculties, faculty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(faculties) == 0 {
		return nil, store.ErrEmptyData
	}

	return faculties, nil
}
