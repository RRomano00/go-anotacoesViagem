package models

import (
	"time"
)

type Note struct {
	Id         int       `json:"id"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Travel_Id  int       `json:"travel_id"`
}

type CreateNoteRequest struct {
	Content  string `json:"content"`
	TravelID int    `json:"travel_id"`
}

type NoteTravel struct {
	Note
	TravelName string
}
