package postgresStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type RoomRepo struct {
	store *Store
}

func (r *RoomRepo) CreateRoom(room *model.Room) error {
	//Validations ???
	return r.store.db.QueryRow("INSERT INTO rooms (number, capacity, free_capacity, hostel_id, room_sex) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		room.Number,
		room.Capacity,
		room.FreeCapacity,
		room.HostelId,
		room.RoomSex,
	).Scan(&room.Id)
}

//Return all roms that located in the same hostelId
func (r *RoomRepo) GetAllRoomsByHostleId(hostelId int) ([]*model.Room, error) {
	rooms := make([]*model.Room, 0)

	rows, err := r.store.db.Query("SELECT id, number, capacity, free_capacity, hostel_id, room_sex FROM rooms WHERE hostel_id = $1", hostelId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		room := &model.Room{}
		if err := rows.Scan(&room.Id, &room.Number, &room.Capacity, &room.FreeCapacity, &room.HostelId, &room.RoomSex); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(rooms) == 0 {
		return nil, store.ErrEmptyData
	}
	return rooms, nil
}
