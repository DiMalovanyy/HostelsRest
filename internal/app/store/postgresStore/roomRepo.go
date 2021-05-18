package postgresStore

import (
	"database/sql"

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

	// log.Print("Hostel id: ", hostelId)
	rows, err := r.store.db.Query("SELECT id, number, capacity, free_capacity, hostel_id, room_sex FROM rooms WHERE hostel_id = $1", hostelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		room := new(model.Room)
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
	// log.Print("log inner: ", len(rooms))
	return rooms, nil
}

func (r *RoomRepo) GetFreeRoomByHostelId(hostelId int) (int, error) {

	rooms, err := r.GetAllRoomsByHostleId(hostelId)
	if err != nil {
		return 0, err
	}

	for _, room := range rooms {
		if room.FreeCapacity > 0 {
			room.FreeCapacity -= 1
			var roomReturnId int
			err := r.store.db.QueryRow("UPDATE rooms SET free_capacity = $1 WHERE id = $2 RETURNING id", room.FreeCapacity, room.Id).Scan(&roomReturnId)
			if err != nil {
				return 0, err
			}
			return roomReturnId, nil
		}

	}

	return 0, store.ErrNoData
}

func (r *RoomRepo) GetRoomByRoomId(roomId int) (*model.Room, error) {
	room := &model.Room{}

	if err := r.store.db.QueryRow(
		"SELECT id, number, capacity, free_capacity, hostel_id, room_sex FROM rooms WHERE id = $1", roomId).Scan(
		&room.Id,
		&room.Number,
		&room.Capacity,
		&room.FreeCapacity,
		&room.HostelId,
		&room.RoomSex,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return room, nil
}
