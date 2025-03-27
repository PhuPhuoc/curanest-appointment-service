package cuspackagecommands

import (
	"context"

	"github.com/google/uuid"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type Commands struct {
	CreateCusPackageAndCusTask *createCusPackageAndTaskHandler
}

type Builder interface {
	BuildCusPackageCmdRepo() CusPackageCommandRepo
	BuildSvcPackageFetcher() SvcPackageFetcher
}

func NewCusPackageCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateCusPackageAndCusTask: NewCreateCusPackageAndTaskHandler(
			b.BuildCusPackageCmdRepo(),
			b.BuildSvcPackageFetcher(),
		),
	}
}

type CusPackageCommandRepo interface {
	CreateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
	CreateCustomizedTasks(ctx context.Context, entities []cuspackagedomain.CustomizedTask) error
}

type SvcPackageFetcher interface {
	GetServicePackageById(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error)
	GetServiceTasksByPackageId(ctx context.Context, svcPackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error)
}

type AppoinmentFetcher interface {
	CreateAppointment(ctx context.Context) error
}

type InvoiceFetcher interface {
	CreateInvoice(ctx context.Context) error
}

type MedicalRecordFetcher interface {
	CreateMedicalRecord(ctx context.Context) error
}
