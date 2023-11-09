package data_responses

type WorkshopsDataListResponses struct {
	Id        int64                `json:"id" db:"id"`
	Title     string               `json:"title" db:"title"`
	StartAt   string               `json:"start_at" db:"start_at"`
	EndAt     string               `json:"end_at" db:"end_at"`
	Workshops []*WorkshopsDataList `json:"workshops"`
}

type WorkshopsDataList struct {
	Id          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	StartAt     string `json:"start_at" db:"start_at"`
	EndAt       string `json:"end_at" db:"end_at"`
}

type WorkshopDetailsResponse struct {
	Id                int64  `json:"id" db:"id"`
	Title             string `json:"title" db:"title"`
	Description       string `json:"description" db:"description"`
	StartAt           string `json:"start_at" db:"start_at"`
	EndAt             string `json:"end_at" db:"end_at"`
	TotalReservations int    `json:"total_reservations" db:"total_reservations"`
}
