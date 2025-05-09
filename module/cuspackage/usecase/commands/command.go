package cuspackagecommands

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type Commands struct {
	CreateCusPackageAndCusTask *createCusPackageAndTaskHandler

	UpdateCustaskStatusDone *updateCustaskStatusDoneHanlder
	UpdateMedicalRecord     *updateMedicalRecordHanlder

	CancelPackage *cancelPackageHanlder
}

type Builder interface {
	BuildTransactionManager() common.TransactionManager
	BuildCusPackageCmdRepo() CusPackageCommandRepo
	BuildSvcPackageFetcher() SvcPackageFetcher
	BuildAppointmentFetcher() AppointmentFetcher
	BuildInvoiceFetcher() InvoiceFetcher
	BuilderPayosConfig() common.PayOSConfig
	BuildExternalGoongAPI() ExternalGoongAPI
	BuildExternalPushNotiService() ExternalPushNotiService
}

func NewCusPackageCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateCusPackageAndCusTask: NewCreateCusPackageAndTaskHandler(
			b.BuildCusPackageCmdRepo(),
			b.BuildSvcPackageFetcher(),
			b.BuildAppointmentFetcher(),
			b.BuildInvoiceFetcher(),
			b.BuildTransactionManager(),
			b.BuilderPayosConfig(),
			b.BuildExternalGoongAPI(),
			b.BuildExternalPushNotiService(),
		),
		UpdateCustaskStatusDone: NewUpdateCustaskStatusDoneHandler(
			b.BuildCusPackageCmdRepo(),
			b.BuildAppointmentFetcher(),
		),
		UpdateMedicalRecord: NewUpdateMedicalRecordHandler(
			b.BuildCusPackageCmdRepo(),
			b.BuildAppointmentFetcher(),
			b.BuildTransactionManager(),
		),
		CancelPackage: NewCancelPackageHandler(
			b.BuildCusPackageCmdRepo(),
			b.BuildAppointmentFetcher(),
			b.BuildTransactionManager(),
			b.BuildExternalPushNotiService(),
		),
	}
}

type CusPackageCommandRepo interface {
	CreateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
	CreateCustomizedTasks(ctx context.Context, entities []cuspackagedomain.CustomizedTask) error
	CreateMedicalRecords(ctx context.Context, entities []cuspackagedomain.MedicalRecord) error

	UpdateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
	UpdateCustomizedTask(ctx context.Context, entity *cuspackagedomain.CustomizedTask) error
	UpdateMedicalRecord(ctx context.Context, entity *cuspackagedomain.MedicalRecord) error

	VerifyAllCusTasksHaveDone(ctx context.Context, cusPackageId uuid.UUID, date time.Time) error
}

type SvcPackageFetcher interface {
	GetServicePackageById(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error)
	GetServiceTasksByPackageId(ctx context.Context, svcPackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error)
}

type AppointmentFetcher interface {
	CreateAppointments(ctx context.Context, entities []appointmentdomain.Appointment) error
	AreNursesAvailable(ctx context.Context, nursingIds []uuid.UUID, dates []time.Time) error

	CheckAppointmentStatusUpcoming(ctx context.Context, cuspackageId uuid.UUID, date time.Time) error
	FindById(ctx context.Context, appointmentId uuid.UUID) (*appointmentdomain.Appointment, error)

	UpdateAppointment(ctx context.Context, entity *appointmentdomain.Appointment) error
	GetAppointmentByCuspackage(ctx context.Context, cuspackageId uuid.UUID) ([]appointmentdomain.Appointment, error)
}

type InvoiceFetcher interface {
	CreateInvoice(ctx context.Context, entity *invoicedomain.Invoice) error
}

type ExternalGoongAPI interface {
	GetGeocodeFromGoong(ctx context.Context, address string) (*GoongAPIResponse, error)
}

type ExternalPushNotiService interface {
	PushNotification(ctx context.Context, req *common.PushNotiRequest) error
}
