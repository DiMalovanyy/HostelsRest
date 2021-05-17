package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type HostelRepo struct {
	Hostels map[int]*model.Hostel
}

func (h *HostelRepo) CreateHostel(hostel *model.Hostel) error {

	//TODD: Validation ???
	hostel.Id = len(h.Hostels) + 1
	h.Hostels[hostel.Id] = hostel

	return nil
}

func (h *HostelRepo) GetHostelsByFucultyId(fucultyId int) ([]*model.Hostel, error) {
	hostels := make([]*model.Hostel, 0)

	for _, hostel := range h.Hostels {
		if hostel.FacultyId == fucultyId {
			hostels = append(hostels, hostel)
		}
	}

	if len(hostels) == 0 {
		return nil, store.ErrEmptyData
	}

	return hostels, nil
}

func (h *HostelRepo) GetHostelById(hostelId int) (*model.Hostel, error) {

	return nil, nil
}
