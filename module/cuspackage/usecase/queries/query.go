package cuspackagequeries

import (
	"context"
	"time"

	"github.com/google/uuid"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type Queries struct {
	FindCusPackageTask    *findCusPackageTaskHandler
	FindCustaskById       *findCustaskByIdHandler
	FindMedicalRecordById *findMedicalRecordByIdHandler
}

type Builder interface {
	BuildCusPackageQueryRepo() CusPackageQueryRepo
}

func NewCusPackageQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindCusPackageTask: NewFindCusPackageTaskDetailHandler(
			b.BuildCusPackageQueryRepo(),
		),
		FindCustaskById: NewFindCustaskByIdHandler(
			b.BuildCusPackageQueryRepo(),
		),
		FindMedicalRecordById: NewFindMedicalRecordByIdHandler(
			b.BuildCusPackageQueryRepo(),
		),
	}
}

type CusPackageQueryRepo interface {
	FindCusPackage(ctx context.Context, id uuid.UUID) (*cuspackagedomain.CustomizedPackage, error)
	FindCusTasks(ctx context.Context, packageId uuid.UUID, estDate time.Time) ([]cuspackagedomain.CustomizedTask, error)
	FindCusTaskById(ctx context.Context, custaskId uuid.UUID) (*cuspackagedomain.CustomizedTask, error)
	FindMedicalRecordById(ctx context.Context, mecicalRecordId uuid.UUID) (*cuspackagedomain.MedicalRecord, error)
}
