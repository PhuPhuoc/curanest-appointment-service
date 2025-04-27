package apppointmentcommands

import (
	"context"

	"github.com/google/uuid"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type Commands struct {
	UpdateApointmentStatus *updateAppointmentHandler
	AssigneNursing         *assignNursingHandler
}

type Builder interface {
	BuildAppointmentCmdRepo() AppointmentCommandRepo
	BuildCusTaskFetcher() CusTaskFetcher
	BuildMedicalRecord() MedicalRecordFetcher
}

func NewAppointmentCmdWithBuilder(b Builder) Commands {
	return Commands{
		UpdateApointmentStatus: NewUpdateAppointmentStatusHandler(
			b.BuildAppointmentCmdRepo(),
			b.BuildCusTaskFetcher(),
			b.BuildMedicalRecord(),
		),
		AssigneNursing: NewAssignNursingHandler(
			b.BuildAppointmentCmdRepo(),
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
	CheckMedicalRecordDone(ctx context.Context, cusPackageId uuid.UUID) error
}
