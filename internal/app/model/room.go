package model

type Room struct {
	Id           int `json:"id"`
	Capacity     int `json:"capacity"`
	FreeCapacity int `json:"free_capacity"`
	HostelId     int `json:"hostel_id"`
	RoomSex      Sex `json:"room_sex"`
}
