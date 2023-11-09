package data_params

type GetWorkshopReservationDataPayload struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
