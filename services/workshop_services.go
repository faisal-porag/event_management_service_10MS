package services

import (
	"event_management_service_10MS/data_responses"
	"event_management_service_10MS/state"
)

func GetWorkshopListForEventServices(
	eventId int64,
	s *state.State,
) (*data_responses.WorkshopsDataListResponses, []*data_responses.WorkshopsDataList, error) {
	event, workshops, err := s.DbRepository.GetEventWithActiveWorkshops(eventId)
	return event, workshops, err
}

func GetWorkshopDetailsServices(
	workshopId int64,
	s *state.State,
) (*data_responses.WorkshopDetailsResponse, error) {
	details, err := s.DbRepository.GetWorkshopDetails(workshopId)
	return details, err
}
