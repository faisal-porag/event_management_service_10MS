package models

import "time"

type Workshop struct {
	Id          int64     `db:"id" json:"id"`
	EventId     int64     `db:"event_id" json:"event_id"`
	StartAt     time.Time `db:"start_at" json:"start_at"`
	EndAt       time.Time `db:"end_at" json:"end_at"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
}
