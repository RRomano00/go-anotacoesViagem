package models

import (
	"time"
)

type Travel struct {
	Id         int
	Title      string
	Start_date time.Time
	End_date   time.Time
}

type CreateTravelRequest struct {
	Title string
}
