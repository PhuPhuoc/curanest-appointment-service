package appointmentrepository

import "github.com/jmoiron/sqlx"

type appointmentRepo struct {
	db *sqlx.DB
}

func NewAppointmentRepo(db *sqlx.DB) *appointmentRepo {
	return &appointmentRepo{
		db: db,
	}
}
