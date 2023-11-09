package services

import (
	"event_management_service_10MS/data_responses"
	"event_management_service_10MS/state"
)

func GetWorkshopReservationServices(
	name string,
	email string,
	s *state.State,
) (*data_responses.WorkshopReservationDataResponse, error) {
	data, err := s.DbRepository.GetWorkshopReservation(name, email)
	return data, err
}
