package models

import "time"

type Event struct {
	ID      int64     `db:"id" json:"id"`
	Title   string    `db:"title" json:"title"`
	StartAt time.Time `db:"start_at" json:"start_at"`
	EndAt   time.Time `db:"end_at" json:"end_at"`
}
