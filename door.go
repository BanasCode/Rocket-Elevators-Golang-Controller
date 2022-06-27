package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_id int, _status string) *Door {
	return &Door{_id, _status}
}
