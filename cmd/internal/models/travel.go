package models

import (
	"time"
)

type Travel struct {
	Id        int
	Title     string
	StartDate time.Time
	EndDate   *time.Time
}

type UpdateTravelRequest struct {
	Title   string
	EndDate *time.Time
}
