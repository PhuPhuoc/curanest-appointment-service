package appointmentqueries

import (
	"context"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/google/uuid"
)

type Queries struct {
	GetAppointment      *getAppointmentsHandler
	FindAppointmentById *findAppointmentByIdHandler
	GetNursingTimeSheet *getNursingTimeSheetHandler
}

type Builder interface {
	BuildAppointmentQueryRepo() AppointmentQueryRepo
}

func NewAppointmentQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetAppointment: NewGetAppointmentsHandler(
			b.BuildAppointmentQueryRepo(),
		),
		FindAppointmentById: NewFindAppointmentByIdHandler(
			b.BuildAppointmentQueryRepo(),
		),
		GetNursingTimeSheet: NewGetNursingTimeSheetHandler(
			b.BuildAppointmentQueryRepo(),
		),
	}
}

type AppointmentQueryRepo interface {
	GetAppointment(ctx context.Context, filter *FilterGetAppointmentDTO) ([]appointmentdomain.Appointment, error)
	FindById(ctx context.Context, appointmentId uuid.UUID) (*appointmentdomain.Appointment, error)
}
