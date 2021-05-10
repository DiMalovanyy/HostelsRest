package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type RoomRepo struct {
	Rooms map[int]*model.Room
}

func (r *RoomRepo) CreateRoom(room *model.Room) error {
	//TODO: Validation ???

	room.Id = len(r.Rooms) + 1
	r.Rooms[room.Id] = room

	return nil
}

//Return all roms that located in the same hostelId as this room
func (r *RoomRepo) GetAllRoomsByHostleId(hostelId int) ([]*model.Room, error) {
	rooms := make([]*model.Room, 0)

	for _, room := range r.Rooms {
		if room.HostelId == hostelId {
			rooms = append(rooms, room)
		}
	}

	if len(rooms) == 0 {
		return nil, store.ErrEmptyData
	}
	return rooms, nil
}
