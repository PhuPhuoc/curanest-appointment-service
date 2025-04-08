package appointmentqueries

import (
	"context"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type Queries struct {
	GetAppointment *getAppointmentsHandler
}

type Builder interface {
	BuildAppointmentQueryRepo() AppointmentQueryRepo
}

func NewAppointmentQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetAppointment: NewGetAppointmentsHandler(
			b.BuildAppointmentQueryRepo(),
		),
	}
}

type AppointmentQueryRepo interface {
	GetAppointment(ctx context.Context, filter *FilterGetAppointmentDTO) ([]appointmentdomain.Appointment, error)
}
