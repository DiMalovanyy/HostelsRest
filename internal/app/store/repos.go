package store

import "github.com/UniverOOP/internal/app/model"

type UserRepo interface {
	CreateUser(*model.User) error
	FindByName(string) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
	Upgrade(int, model.Sex, int, int) error

	// GetUsersByRoomId() ([]*model.User, error)
}

type HostelRepo interface {
	CreateHostel(*model.Hostel) error
	GetHostelsByFucultyId(int) ([]*model.Hostel, error)
}

type FacultyRepo interface {
	CreateFaculty(*model.Faculty) error
	GetAllFaculties() ([]*model.Faculty, error)
}

type RoomRepo interface {
	CreateRoom(*model.Room) error
	//Return all roms that located in the same hostelId as this room
	GetAllNeighborlyRooms() ([]*model.Room, error)
}
