package store

import "github.com/UniverOOP/internal/app/model"

type UserRepo interface {
	CreateUser(*model.User) error
	FindByName(string) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindById(int) (*model.User, error)
	Upgrade(int, model.Sex, int, int, int) error

	GetAllUsersByRoomId(int) ([]*model.User, error)
	// GetUsersByRoomId() ([]*model.User, error)
}

type HostelRepo interface {
	CreateHostel(*model.Hostel) error
	GetHostelsByFucultyId(int) ([]*model.Hostel, error)
	// GetHostelByName(string) (*model.Hostel, error)
	GetHostelById(int) (*model.Hostel, error)
}

type FacultyRepo interface {
	CreateFaculty(*model.Faculty) error
	GetAllFaculties() ([]*model.Faculty, error)
	GetFacultyByName(string) (*model.Faculty, error)
}

type RoomRepo interface {
	CreateRoom(*model.Room) error
	//Return all roms that located in the same hostelId
	GetAllRoomsByHostleId(int) ([]*model.Room, error)
	GetFreeRoomByHostelId(int) (int, error)
	GetRoomByRoomId(int) (*model.Room, error)
}
