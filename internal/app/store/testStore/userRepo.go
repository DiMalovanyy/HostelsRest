package testStore

import (
	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
)

type UserRepo struct {
	userById map[int]*model.User
}

func (repo *UserRepo) CreateUser(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	user.Id = len(repo.userById) + 1
	user.Sex = model.UNDEFINED
	user.RoomId = 0
	user.FacultyId = 0
	repo.userById[user.Id] = user
	return nil
}

func (repo *UserRepo) FindByName(name string) (*model.User, error) {
	for _, u := range repo.userById {
		if u.Name == name {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	for _, u := range repo.userById {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (repo *UserRepo) FindById(id int) (*model.User, error) {
	u, ok := repo.userById[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (repo *UserRepo) Upgrade(userId int, sex model.Sex, roomId int, facultyId int) error {
	//Todo validation

	u, err := repo.FindById(userId)
	if err != nil {
		return err
	}

	u.Sex = sex
	u.RoomId = roomId
	u.FacultyId = facultyId

	return nil
}
