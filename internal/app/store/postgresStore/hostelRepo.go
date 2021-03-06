package postgresStore

import (
	"database/sql"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type HostelRepo struct {
	store *Store
}

func (h *HostelRepo) CreateHostel(hostel *model.Hostel) error {

	//TODO: Validation
	return h.store.db.QueryRow("INSERT INTO hostels (description, faculty_id) VALUES ($1, $2) RETURNING id",
		hostel.Description,
		hostel.FacultyId,
	).Scan(&hostel.Id)
}

func (h *HostelRepo) GetHostelsByFucultyId(fucultyId int) ([]*model.Hostel, error) {

	hostels := make([]*model.Hostel, 0)

	rows, err := h.store.db.Query("SELECT id, description, faculty_id FROM hostels WHERE faculty_id = $1", fucultyId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		hostel := &model.Hostel{}
		if err := rows.Scan(&hostel.Id, &hostel.Description, &hostel.FacultyId); err != nil {
			return nil, err
		}
		hostels = append(hostels, hostel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(hostels) == 0 {
		return nil, store.ErrEmptyData
	}
	return hostels, nil
}

func (h *HostelRepo) GetHostelById(hostelId int) (*model.Hostel, error) {
	hostel := &model.Hostel{}

	if err := h.store.db.QueryRow(
		"SELECT id, description, faculty_id FROM hostels WHERE id = $1", hostelId,
	).Scan(
		&hostel.Id,
		&hostel.Description,
		&hostel.FacultyId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return hostel, nil
}
