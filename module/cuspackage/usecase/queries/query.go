package cuspackagequeries

import (
	"context"
	"time"

	"github.com/google/uuid"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type Queries struct {
	FindCusPackageTask *findCusPackageTaskHandler
}

type Builder interface {
	BuildCusPackageQueryRepo() CusPackageQueryRepo
}

func NewCusPackageQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindCusPackageTask: NewFindCusPackageTaskDetailHandler(
			b.BuildCusPackageQueryRepo(),
		),
	}
}

type CusPackageQueryRepo interface {
	FindCusPackage(ctx context.Context, id uuid.UUID) (*cuspackagedomain.CustomizedPackage, error)
	FindCusTasks(ctx context.Context, packageId uuid.UUID, estDate time.Time) ([]cuspackagedomain.CustomizedTask, error)
}
