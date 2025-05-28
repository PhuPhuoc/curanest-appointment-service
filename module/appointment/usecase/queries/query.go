package appointmentqueries

import (
	"context"
	"time"

	"github.com/google/uuid"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type Queries struct {
	GetAppointment      *getAppointmentsHandler
	FindAppointmentById *findAppointmentByIdHandler
	GetNursingTimeSheet *getNursingTimeSheetHandler
	GetNursingAvailable *getNursingAvailableHandler

	CheckNursesAvailability *checkNursesAvailabilityHandler
	GetDashboardData        *getItemDashboardHandler
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
		GetDashboardData: NewGetItemDashboardHandler(
			b.BuildAppointmentQueryRepo(),
		),
	}
}

type AppointmentQueryRepo interface {
	GetAppointment(ctx context.Context, filter *FilterGetAppointmentDTO) ([]appointmentdomain.Appointment, error)
	FindById(ctx context.Context, appointmentId uuid.UUID) (*appointmentdomain.Appointment, error)
	GetAppointmentInDate(ctx context.Context, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error)
	GetAppointmentInADayOfNursing(ctx context.Context, nursingId uuid.UUID, estStartDate, estEndDate time.Time) ([]appointmentdomain.Appointment, error)

	IsNurseAvailability(ctx context.Context, nursingIds uuid.UUID, startDate, endDate time.Time) error

	GetDashboardData(ctx context.Context, filter *FilterDashboardDTO) (int, int, int, error)
}

type CusPackageFetcher interface {
	GetCusPackageByIds(ctx context.Context, cuspackageId []uuid.UUID) (map[uuid.UUID]cuspackagedomain.CustomizedPackage, error)
}

type ExternalNursingService interface {
	GetNursingByServiceIdRPC(ctx context.Context, serviceId uuid.UUID) ([]NurseDTO, error)
}
