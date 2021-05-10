package postgresStore

import (
	"testing"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestRoomRepo_CreateRoom(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("rooms", "faculty", "hostels")

	s := New(db)
	facRepo := s.Faculty()
	hostelRepo := s.Hostel()
	roomRepo := s.Room()

	h := model.TestHostel(t)
	fac := model.TestFaculty(t)
	room := model.TestRoom(t)

	facRepo.CreateFaculty(fac)
	fac, _ = facRepo.GetFacultyByName(fac.Name)
	h.FacultyId = fac.Id

	hostelRepo.CreateHostel(h)

	room.HostelId = h.Id
	assert.NoError(t, roomRepo.CreateRoom(room))
	assert.NotNil(t, room)
}

func TestRoomRepo_GetRoomsByHostelId(t *testing.T) {
	db, teardown := TestDB(t, databaseURL)
	defer teardown("rooms", "faculty", "hostels")

	s := New(db)
	facRepo := s.Faculty()
	hostelRepo := s.Hostel()
	roomRepo := s.Room()

	h := model.TestHostel(t)
	fac := model.TestFaculty(t)
	room := model.TestRoom(t)

	facRepo.CreateFaculty(fac)
	fac, _ = facRepo.GetFacultyByName(fac.Name)
	h.FacultyId = fac.Id

	hostelRepo.CreateHostel(h)
	room.HostelId = h.Id

	roomRepo.CreateRoom(room)

	_, err := roomRepo.GetAllRoomsByHostleId(10)
	assert.EqualError(t, err, store.ErrEmptyData.Error())

	rs, err := roomRepo.GetAllRoomsByHostleId(h.Id)
	assert.NoError(t, err)
	assert.NotEqual(t, len(rs), 0)
	assert.Equal(t, rs[0].Number, room.Number)

	hs, err := hostelRepo.GetHostelsByFucultyId(fac.Id)
	assert.NotNil(t, rs)
	assert.NoError(t, err)
	assert.NotEqual(t, len(hs), 0)
	assert.Equal(t, hs[0].Description, h.Description)
}
