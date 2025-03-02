package servicecommands

import (
	"context"

	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

type Commands struct {
	CreateService *createServiceHandler
}

type Builder interface {
	BuildServiceCmdRepo() ServiceCommandRepo
}

func NewServiceCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateService: NewCreateServiceHandler(
			b.BuildServiceCmdRepo(),
		),
	}
}

type ServiceCommandRepo interface {
	Create(ctx context.Context, entity *servicedomain.Service) error
}
