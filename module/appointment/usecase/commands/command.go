package apppointmentcommands

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type Commands struct {
	UpdateApointmentStatus *updateAppointmentHandler
	UpdateStatusUpcoming   *updateStatusUpcomingHandler
	AssigneNursing         *assignNursingHandler
}

type Builder interface {
	BuildTransactionManager() common.TransactionManager
	BuildAppointmentCmdRepo() AppointmentCommandRepo
	BuildCusTaskFetcher() CusTaskFetcher
	BuildMedicalRecord() MedicalRecordFetcher
	BuildExternalGoongAPI() ExternalGoongAPI
}

func NewAppointmentCmdWithBuilder(b Builder) Commands {
	return Commands{
		UpdateApointmentStatus: NewUpdateAppointmentStatusHandler(
			b.BuildAppointmentCmdRepo(),
			b.BuildCusTaskFetcher(),
			b.BuildMedicalRecord(),
		),
		UpdateStatusUpcoming: NewUpdateStatusUpcomingHandler(
			b.BuildAppointmentCmdRepo(),
			b.BuildExternalGoongAPI(),
		),
		AssigneNursing: NewAssignNursingHandler(
			b.BuildAppointmentCmdRepo(),
			b.BuildTransactionManager(),
			b.BuildMedicalRecord(),
		),
	}
}

type AppointmentCommandRepo interface {
	UpdateAppointment(ctx context.Context, entity *appointmentdomain.Appointment) error
}

type CusTaskFetcher interface {
	CheckCusTasksAllDone(ctx context.Context, cusPackageId uuid.UUID) error
}

type MedicalRecordFetcher interface {
	FindMedicalRecordByAppsId(ctx context.Context, appsId uuid.UUID) (*cuspackagedomain.MedicalRecord, error)
	UpdateMedicalRecord(ctx context.Context, entity *cuspackagedomain.MedicalRecord) error

	CheckMedicalRecordDone(ctx context.Context, cusPackageId uuid.UUID) error
}

type ExternalGoongAPI interface {
	GetDistanceFromGoong(ctx context.Context, originCode, destinationCode string) (*DistanceMatrixResponse, error)
}
