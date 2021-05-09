package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type Sex int

const (
	MEN Sex = iota
	WOMEN
	UNDEFINED
)

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`

	//Not main fields
	Sex       Sex `json:"-"`
	RoomId    int `json:"-"`
	FacultyId int `json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 50)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encrypteString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}

	return nil
}

func encrypteString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

//TODO: Test it
func (u *User) IsUserUpdated() bool {
	return !(u.Sex == UNDEFINED || u.RoomId == 0 || u.FacultyId == 0)
}
