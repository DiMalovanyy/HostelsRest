package store

type Store interface {
	//Repopsitories

	User() UserRepo
	Faculty() FacultyRepo
}
