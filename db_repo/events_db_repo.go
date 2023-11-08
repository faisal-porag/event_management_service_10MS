package db_repo

import (
	"event_management_service_10MS/data_responses"
)

func (pr *MySQLRepository) GetAllMenuTypeListInfo(
	currentPage int,
	itemsPerPage int,
) ([]*data_responses.EventListResponses, *data_responses.Pagination, error) {
	offset := (currentPage - 1) * itemsPerPage
	var dataList []*data_responses.EventListResponses

	query := `
        SELECT id, title, start_at, end_at, (
            SELECT COUNT(*) FROM events WHERE start_at > NOW() OR end_at > NOW()
        ) as total_count
        FROM events
        WHERE start_at > NOW() OR end_at > NOW()
        LIMIT ? OFFSET ?
    `

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
