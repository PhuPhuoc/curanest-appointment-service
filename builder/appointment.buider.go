package builder

import (
	"github.com/jmoiron/sqlx"

	appointmentrepository "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/infars/repository"
	apppointmentcommands "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/commands"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

type builderOfAppointment struct {
	db *sqlx.DB
}

func NewAppointmentBuilder(db *sqlx.DB) builderOfAppointment {
	return builderOfAppointment{db: db}
}

func (s builderOfAppointment) BuildAppointmentCmdRepo() apppointmentcommands.AppointmentCommandRepo {
	return appointmentrepository.NewAppointmentRepo(s.db)
}

func (s builderOfAppointment) BuildAppointmentQueryRepo() appointmentqueries.AppointmentQueryRepo {
	return appointmentrepository.NewAppointmentRepo(s.db)
}
