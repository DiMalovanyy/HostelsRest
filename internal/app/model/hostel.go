package model

type Hostel struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	FacultyId   int    `json:"faculty_id"`
}
