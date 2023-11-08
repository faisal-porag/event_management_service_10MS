package db_repo

import (
	"errors"
	"event_management_service_10MS/data_responses"
)

var NoDataFound = errors.New("no data found")
var RowsPerPageZeroError = errors.New("per_page value must be grater than zero")

func GetPaginationInfoData(totalCount int64, CurrentPage, limitPerPage int) (*data_responses.Pagination, error) {
	// Calculate the total number of pages
	if limitPerPage == 0 {
		return nil, RowsPerPageZeroError
	}
	totalPages := totalCount / int64(limitPerPage)
	if totalCount%int64(limitPerPage) > 0 {
		totalPages++
	}

	hasNextPage := int64(CurrentPage) < totalPages
	pagingData := &data_responses.Pagination{
		Total:       totalCount,
		PerPage:     limitPerPage,
		CurrentPage: CurrentPage,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
	}

	return pagingData, nil
}
