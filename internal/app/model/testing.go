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
