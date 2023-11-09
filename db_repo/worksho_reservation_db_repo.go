package db_repo

import (
	"database/sql"
	"event_management_service_10MS/data_responses"
)

func (pr *MySQLRepository) GetWorkshopReservation(
	name string,
	email string,
) (*data_responses.WorkshopReservationDataResponse, error) {
	var result data_responses.WorkshopReservationData
	query := `
		SELECT
			reservations.id as reservation_id,
			reservations.name as reservation_name,
			reservations.email as reservation_email,
			events.id as event_id,
			events.title as event_title,
			events.start_at as event_start_at,
			events.end_at as event_end_at,
			workshops.id as workshop_id,
			workshops.title as workshop_title,
			workshops.description as workshop_description,
			workshops.start_at as workshop_start_at,
			workshops.end_at as workshop_end_at
		FROM
			reservations
		JOIN workshops ON reservations.workshop_id = workshops.id
		JOIN events ON workshops.event_id = events.id
		WHERE reservations.name = ? AND reservations.email = ?`

	err := pr.db.Get(&result, query, name, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, NoDataFound
		} else {
			return nil, err
		}
	}

	finalData := data_responses.WorkshopReservationDataResponse{
		Reservation: struct {
			Id    int64  `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			Id:    result.ReservationId,
			Name:  result.ReservationName,
			Email: result.ReservationEmail,
		},
		Event: struct {
			Id      int64  `json:"id"`
			Title   string `json:"title"`
			StartAt string `json:"start_at"`
			EndAt   string `json:"end_at"`
		}{
			Id:      result.EventId,
			Title:   result.EventTitle,
			StartAt: result.EventStartAt,
			EndAt:   result.EventEndAt,
		},
		Workshop: struct {
			Id          int64  `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			StartAt     string `json:"start_at"`
			EndAt       string `json:"end_at"`
		}{
			Id:          result.WorkshopId,
			Title:       result.WorkshopTitle,
			Description: result.WorkshopDescription,
			StartAt:     result.WorkshopStartAt,
			EndAt:       result.WorkshopEndAt,
		},
	}

	return &finalData, err
}
