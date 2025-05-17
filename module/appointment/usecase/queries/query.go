package appointmentqueries

import (
	"context"
	"time"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	"github.com/google/uuid"
)

type Queries struct {
	GetAppointment      *getAppointmentsHandler
	FindAppointmentById *findAppointmentByIdHandler
	GetNursingTimeSheet *getNursingTimeSheetHandler
	GetNursingAvailable *getNursingAvailableHandler

	CheckNursesAvailability *checkNursesAvailabilityHandler
}

type Builder interface {
	BuildAppointmentQueryRepo() AppointmentQueryRepo
	BuildExternalNurseServiceApiQuery() ExternalNursingService
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
		CheckNursesAvailability: NewCheckNursesAvailabilityHandler(
			b.BuildAppointmentQueryRepo(),
		),
		GetNursingAvailable: NewGetNursingAvailableHandler(
			b.BuildAppointmentQueryRepo(),
			b.BuildExternalNurseServiceApiQuery(),
		),
	}
}

type AppointmentQueryRepo interface {
	GetAppointment(ctx context.Context, filter *FilterGetAppointmentDTO) ([]appointmentdomain.Appointment, error)
	FindById(ctx context.Context, appointmentId uuid.UUID) (*appointmentdomain.Appointment, error)
	GetAppointmentInDate(ctx context.Context, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error)
	GetAppointmentInADayOfNursing(ctx context.Context, nursingId uuid.UUID, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error)

	IsNurseAvailability(ctx context.Context, nursingIds uuid.UUID, startDate, endDate time.Time) error
}

type ExternalNursingService interface {
	GetNursingByServiceIdRPC(ctx context.Context, serviceId uuid.UUID) ([]NurseDTO, error)
}
