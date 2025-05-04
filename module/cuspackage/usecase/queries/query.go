package cuspackagequeries

import (
	"context"
	"time"

	"github.com/google/uuid"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type Queries struct {
	FindCusPackageTask        *findCusPackageTaskHandler
	FindCuspackageById        *findCusPackageByIdHandler
	FindCustaskById           *findCustaskByIdHandler
	FindMedicalRecordById     *findMedicalRecordByIdHandler
	FindMedicalRecordByAppsId *findMedicalRecordByAppsIdHandler
}

type Builder interface {
	BuildCusPackageQueryRepo() CusPackageQueryRepo
}

func NewCusPackageQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindCusPackageTask: NewFindCusPackageTaskDetailHandler(
			b.BuildCusPackageQueryRepo(),
		),
		FindCuspackageById: NewFindCusPackageByIdHandler(
			b.BuildCusPackageQueryRepo(),
		),
		FindCustaskById: NewFindCustaskByIdHandler(
			b.BuildCusPackageQueryRepo(),
		),
		FindMedicalRecordById: NewFindMedicalRecordByIdHandler(
			b.BuildCusPackageQueryRepo(),
		),
		FindMedicalRecordByAppsId: NewFindMedicalRecordByAppsIdHandler(
			b.BuildCusPackageQueryRepo(),
		),
	}
}

type CusPackageQueryRepo interface {
	FindCusPackage(ctx context.Context, id uuid.UUID) (*cuspackagedomain.CustomizedPackage, error)
	FindCusTasks(ctx context.Context, packageId uuid.UUID, estDate time.Time) ([]cuspackagedomain.CustomizedTask, error)
	FindCusTaskById(ctx context.Context, custaskId uuid.UUID) (*cuspackagedomain.CustomizedTask, error)
	FindMedicalRecordById(ctx context.Context, mecicalRecordId uuid.UUID) (*cuspackagedomain.MedicalRecord, error)
	FindMedicalRecordByAppsId(ctx context.Context, appsId uuid.UUID) (*cuspackagedomain.MedicalRecord, error)
}
