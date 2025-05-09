package servicecommands

import (
	"context"

	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

type Commands struct {
	CreateService *createServiceHandler
	UpdateService *updateServiceHandler
}

type Builder interface {
	BuildServiceCmdRepo() ServiceCommandRepo
}

func NewServiceCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateService: NewCreateServiceHandler(
			b.BuildServiceCmdRepo(),
		),
		UpdateService: NewUpdateServiceHandler(
			b.BuildServiceCmdRepo(),
		),
	}
}

type ServiceCommandRepo interface {
	Create(ctx context.Context, entity *servicedomain.Service) error
	Update(ctx context.Context, entity *servicedomain.Service) error
}
