package db_repo

import (
	"database/sql"
	"event_management_service_10MS/data_responses"
)

func (pr *MySQLRepository) GetAllEventListInfo(
	currentPage int,
	itemsPerPage int,
) ([]*data_responses.EventListResponses, *data_responses.Pagination, error) {
	offset := (currentPage - 1) * itemsPerPage
	var dataList []*data_responses.EventListResponses

	query := `
	   SELECT id, title, start_at, end_at, (
	       SELECT COUNT(*) FROM events WHERE start_at > NOW()
	   ) as total_count
	   FROM events
	   WHERE start_at > NOW()
	   LIMIT ? OFFSET ?
	`

	//query := `
	//    SELECT id, title, start_at, end_at, (
	//        SELECT COUNT(*) FROM events WHERE NOW() BETWEEN start_at AND end_at
	//    ) as total_count
	//    FROM events
	//    WHERE NOW() BETWEEN start_at AND end_at
	//    LIMIT ? OFFSET ?
	//`

	var totalCount int64
	err := pr.db.Select(&dataList, query, itemsPerPage, offset)
	if err != nil {
		return nil, nil, err
	}

	if len(dataList) > 0 {
		totalCount = dataList[0].TotalCount
	}

	pagingData, err := GetPaginationInfoData(totalCount, currentPage, itemsPerPage)

	return dataList, pagingData, err
}

func (pr *MySQLRepository) GetEventDetailsByEventId(eventId int64) (*data_responses.EventDetailsResponses, int, error) {
	event := data_responses.EventDetailsResponses{}
	err := pr.db.Get(&event, "SELECT id, title, start_at, end_at FROM events WHERE id = ?", eventId)
	if err == sql.ErrNoRows {
		return nil, 0, NoDataFound
	} else if err != nil {
		return nil, 0, err
	}

	var totalWorkshops int
	err = pr.db.Get(&totalWorkshops, "SELECT COUNT(*) FROM workshops WHERE event_id = ?", eventId)
	if err != nil {
		return nil, 0, err
	}

	return &event, totalWorkshops, err
}
