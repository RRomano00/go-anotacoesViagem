package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id         uuid.UUID
	Content    string
	Created_at time.Time
}
