package models

import (
	"time"
)

type Travel struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type UpdateTravelRequest struct {
	Title   string
	EndDate *time.Time
}
