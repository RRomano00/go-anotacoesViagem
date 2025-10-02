package models

import (
	"time"

	"github.com/google/uuid"
)

type Travel struct {
	id         uuid.UUID
	title      string
	start_date time.Time
	end_date   time.Time
}
