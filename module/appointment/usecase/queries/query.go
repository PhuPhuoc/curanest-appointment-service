package appointmentqueries

import (
	"context"
	"time"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
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
	BuilderCusPackageFetcher() CusPackageFetcher
}

func NewAppointmentQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetAppointment: NewGetAppointmentsHandler(
			b.BuildAppointmentQueryRepo(),
			b.BuilderCusPackageFetcher(),
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

type CusPackageFetcher interface {
	GetCusPackageByIds(ctx context.Context, cuspackageId []uuid.UUID) (map[uuid.UUID]cuspackagedomain.CustomizedPackage, error)
}

type ExternalNursingService interface {
	GetNursingByServiceIdRPC(ctx context.Context, serviceId uuid.UUID) ([]NurseDTO, error)
}
