package models

type Reservation struct {
	Id         int64  `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Email      string `db:"email" json:"email"`
	WorkshopId int64  `db:"workshop_id" json:"workshop_id"`
}
