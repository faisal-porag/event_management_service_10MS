package services

import (
	"event_management_service_10MS/data_responses"
	"event_management_service_10MS/state"
)

func GetEventsServices(
	currentPage int,
	itemsPerPage int,
	s *state.State,
) ([]*data_responses.EventListResponses, *data_responses.Pagination, error) {
	data, pageInfo, err := s.DbRepository.GetAllEventListInfo(currentPage, itemsPerPage)
	return data, pageInfo, err
}

func GetEventDetailsByEventIdServices(
	eventId int64,
	s *state.State,
) (*data_responses.EventDetailsResponses, error) {
	data, totalWorkshops, err := s.DbRepository.GetEventDetailsByEventId(eventId)
	if data != nil {
		data.TotalWorkshops = totalWorkshops
	}

	return data, err
}
