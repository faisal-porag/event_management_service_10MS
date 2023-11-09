package db_repo

import (
	"database/sql"
	"event_management_service_10MS/data_responses"
	"github.com/rs/zerolog/log"
)

func (pr *MySQLRepository) GetActiveWorkshopsForEvent(eventId int64) ([]*data_responses.WorkshopsDataList, error) {
	var workshops []*data_responses.WorkshopsDataList
	query := `
		SELECT id, title, description, start_at, end_at
		FROM workshops
		WHERE event_id = ? AND start_at > NOW()
	`

	err := pr.db.Select(&workshops, query, eventId)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return workshops, nil
}

func (pr *MySQLRepository) GetEventWithActiveWorkshops(
	eventId int64,
) (*data_responses.WorkshopsDataListResponses, []*data_responses.WorkshopsDataList, error) {
	var event data_responses.WorkshopsDataListResponses
	err := pr.db.Get(&event, `SELECT id, title, start_at, end_at FROM events WHERE id = ?`, eventId)

	if err == sql.ErrNoRows {
		return nil, nil, NoDataFound
	} else if err != nil {
		return nil, nil, err
	}

	workshops, err := pr.GetActiveWorkshopsForEvent(eventId)
	if err != nil {
		return nil, nil, err
	}

	return &event, workshops, nil
}

func (pr *MySQLRepository) GetWorkshopDetails(
	workshopId int64,
) (*data_responses.WorkshopDetailsResponse, error) {
	var workshop data_responses.WorkshopDetailsResponse
	err := pr.db.Get(
		&workshop,
		`SELECT id, title, description, start_at, end_at, 
			 	(SELECT COUNT(*) FROM reservations WHERE workshop_id = ?) as total_reservations
			FROM workshops WHERE id = ?`,
		workshopId, workshopId,
	)
	if err == sql.ErrNoRows {
		return nil, NoDataFound
	} else if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return &workshop, err
}
