package data_responses

type EventListResponses struct {
	Id         int64  `db:"id" json:"id"`
	Title      string `db:"title" json:"title"`
	StartAt    string `db:"start_at" json:"start_at"`
	EndAt      string `db:"end_at" json:"end_at"`
	TotalCount int64  `db:"total_count" json:"-"`
}

type EventDetailsResponses struct {
	Id             int64  `db:"id" json:"id"`
	Title          string `db:"title" json:"title"`
	StartAt        string `db:"start_at" json:"start_at"`
	EndAt          string `db:"end_at" json:"end_at"`
	TotalWorkshops int    `json:"total_workshops"`
}
