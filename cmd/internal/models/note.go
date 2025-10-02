package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	id         uuid.UUID
	content    string
	created_at time.Time
}
