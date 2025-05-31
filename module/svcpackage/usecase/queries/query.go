package svcpackagequeries

import (
	"context"

	"github.com/google/uuid"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type Queries struct {
	GetServicePackages *getSvcPackagesHandler
	GetServiceTasks    *getSvcTasksHandler

	GetSvcPackageUsageCount *getSvcPackagesUsageCountHandler
}

type Builder interface {
	BuildSvcPackageQueryRepo() SvcPackageQueryRepo
}

func NewSvcPackageQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetServicePackages: NewGetServicePackagesHandler(
			b.BuildSvcPackageQueryRepo(),
		),
		GetServiceTasks: NewGetServiceTasksHandler(
			b.BuildSvcPackageQueryRepo(),
		),
		GetSvcPackageUsageCount: NewGetServicePackagesUsageCountHandler(
			b.BuildSvcPackageQueryRepo(),
		),
	}
}

type SvcPackageQueryRepo interface {
	GetSvcPackges(ctx context.Context, serviceId uuid.UUID) ([]svcpackagedomain.ServicePackage, error)
	GetSvcTasks(ctx context.Context, svcpackageId uuid.UUID) ([]svcpackagedomain.ServiceTask, error)

	GetSvcPackageUsageCount(ctx context.Context, cateId uuid.UUID) ([]svcpackagedomain.ServicePackageUsage, error)
}
