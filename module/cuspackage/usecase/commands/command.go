package cuspackagecommands

import (
	"context"

	"github.com/google/uuid"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type Commands struct{}

type Builder interface {
	BuildCusPackageCmdRepo() CusPackageCommandRepo
}

func NewCusPackageCmdWithBuilder(b Builder) Commands {
	return Commands{}
}

type CusPackageCommandRepo interface {
	// CreateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
	// CreateCustomizedTask(ctx context.Context, entity *cuspackagedomain.CustomizedTask) error
}

type SvcPackageFetcher interface {
	GetServicePackageById(ctx context.Context, svcPackageId uuid.UUID) (*svcpackagedomain.ServicePackage, error)
	GetServiceTaskById(ctx context.Context, svcTaskId uuid.UUID) (*svcpackagedomain.ServiceTask, error)
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
