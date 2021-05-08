package store

type Store interface {
	//Repopsitories

	User() UserRepo
}
