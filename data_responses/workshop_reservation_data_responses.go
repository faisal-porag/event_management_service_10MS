package data_responses

type WorkshopReservationDataResponse struct {
	Reservation struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"reservation"`
	Event struct {
		Id      int64  `json:"id"`
		Title   string `json:"title"`
		StartAt string `json:"start_at"`
		EndAt   string `json:"end_at"`
	} `json:"event"`
	Workshop struct {
		Id          int64  `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		StartAt     string `json:"start_at"`
		EndAt       string `json:"end_at"`
	} `json:"workshop"`
}

type WorkshopReservationData struct {
	ReservationId       int64  `db:"reservation_id"`
	ReservationName     string `db:"reservation_name"`
	ReservationEmail    string `db:"reservation_email"`
	EventId             int64  `db:"event_id"`
	EventTitle          string `db:"event_title"`
	EventStartAt        string `db:"event_start_at"`
	EventEndAt          string `db:"event_end_at"`
	WorkshopId          int64  `db:"workshop_id"`
	WorkshopTitle       string `db:"workshop_title"`
	WorkshopDescription string `db:"workshop_description"`
	WorkshopStartAt     string `db:"workshop_start_at"`
	WorkshopEndAt       string `db:"workshop_end_at"`
}
