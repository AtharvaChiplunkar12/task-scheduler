package models

import (
	"time"

	"github.com/jackc/pgtype"
)

type UserRequest struct {
	Task string `json:"task"`
	ScheduledAt time.Time `json:"scheduled_at"`
}

type Task struct {
	Id          string
	Task     string
	ScheduledAt pgtype.Timestamp
	PickedAt    pgtype.Timestamp
	StartedAt   pgtype.Timestamp
	CompletedAt pgtype.Timestamp
	FailedAt    pgtype.Timestamp
}

