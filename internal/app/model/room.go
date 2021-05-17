package model

type Room struct {
	Id           int    `json:"id"`
	Number       string `json:"number"`
	Capacity     int    `json:"capacity"`
	FreeCapacity int    `json:"free_capacity"`
	HostelId     int    `json:"hostel_id"`
	RoomSex      string `json:"room_sex"`
}
