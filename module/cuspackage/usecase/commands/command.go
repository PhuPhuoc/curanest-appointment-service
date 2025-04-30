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
}

type Builder interface {
	BuildTransactionManager() common.TransactionManager
	BuildCusPackageCmdRepo() CusPackageCommandRepo
	BuildSvcPackageFetcher() SvcPackageFetcher
	BuildAppointmentFetcher() AppointmentFetcher
	BuildInvoiceFetcher() InvoiceFetcher
	BuilderPayosConfig() common.PayOSConfig
	BuildExternalGoongAPI() ExternalGoongAPI
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
		),
		UpdateCustaskStatusDone: NewUpdateCustaskStatusDoneHandler(
			b.BuildCusPackageCmdRepo(),
			b.BuildAppointmentFetcher(),
		),
	}
}

type CusPackageCommandRepo interface {
	CreateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
	CreateCustomizedTasks(ctx context.Context, entities []cuspackagedomain.CustomizedTask) error
	CreateMedicalRecords(ctx context.Context, entities []cuspackagedomain.MedicalRecord) error

	UpdateCustomizedTask(ctx context.Context, entity *cuspackagedomain.CustomizedTask) error
}

type SvcPackageFetcher interface {
	GetServicePackageById(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error)
	GetServiceTasksByPackageId(ctx context.Context, svcPackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error)
}

type AppointmentFetcher interface {
	CreateAppointments(ctx context.Context, entities []appointmentdomain.Appointment) error
	AreNursesAvailable(ctx context.Context, nursingIds []uuid.UUID, dates []time.Time) error

	CheckAppointmentStatusUpcoming(ctx context.Context, cuspackageId uuid.UUID, date time.Time) error
}

type InvoiceFetcher interface {
	CreateInvoice(ctx context.Context, entity *invoicedomain.Invoice) error
}

type ExternalGoongAPI interface {
	GetGeocodeFromGoong(ctx context.Context, address string) (*GoongAPIResponse, error)
}
