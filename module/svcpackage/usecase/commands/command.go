package svcpackagecommands

import (
	"context"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type Commands struct {
	CreatePackage *createSvcPackageHandler
	UpdatePackage *updateSvcPackageHandler

	CreateTask      *createSvcTaskHandler
	UpdateTask      *updateSvcTaskHandler
	UpdateTaskOrder *updateTaskOrderHandler
}

type Builder interface {
	BuildSvcPackageCmdRepo() SvcPackageCommandRepo
}

func NewSvcPackageCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreatePackage: NewCreateSvcPackageHandler(
			b.BuildSvcPackageCmdRepo(),
		),
		UpdatePackage: NewUpdateSvcPackageHandler(
			b.BuildSvcPackageCmdRepo(),
		),
		CreateTask: NewCreateSvcTaskHandler(
			b.BuildSvcPackageCmdRepo(),
		),
		UpdateTask: NewUpdateSvcTaskHandler(
			b.BuildSvcPackageCmdRepo(),
		),
		UpdateTaskOrder: NewUpdateTaskOrderHandler(
			b.BuildSvcPackageCmdRepo(),
		),
	}
}

type SvcPackageCommandRepo interface {
	CreatePackage(ctx context.Context, entity *svcpackagedomain.ServicePackage) error
	UpdatePackage(ctx context.Context, entity *svcpackagedomain.ServicePackage) error
	CreateTask(ctx context.Context, entity *svcpackagedomain.ServiceTask) error
	UpdateTask(ctx context.Context, entity *svcpackagedomain.ServiceTask) error
	UpdateTaskOrder(ctx context.Context, entity []svcpackagedomain.ServiceTask) error
}
