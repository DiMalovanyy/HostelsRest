package model

import "testing"

func TestUser(t *testing.T) *User {

	return &User{
		Email:    "example@gmail.com",
		Name:     "Test Name",
		Password: "password",
	}
}

func TestFaculty(t *testing.T) *Faculty {
	return &Faculty{
		Name: "Cybernetic",
	}
}

func TestHostel(t *testing.T) *Hostel {
	return &Hostel{
		Description: "Hostel #1",
		// FacultyId:  ,
	}
}
