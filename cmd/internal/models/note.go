package models

import (
	"time"
)

type Note struct {
	Id         int
	Content    string
	Created_at time.Time
	Travel_Id  int
}

type CreateNoteRequest struct {
	Content  string
	TravelID int
}

type NoteTravel struct {
	Note
	TravelName string
}
