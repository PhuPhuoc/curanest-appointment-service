package svcpackagecommands

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type createSvcPackageHandler struct {
	cmdRepo SvcPackageCommandRepo
}

func NewCreateSvcPackageHandler(cmdRepo SvcPackageCommandRepo) *createSvcPackageHandler {
	return &createSvcPackageHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createSvcPackageHandler) Handle(ctx context.Context, serviceId uuid.UUID, dto *ServicePackageDTO) error {
	svcId := common.GenUUID()

	entity, _ := svcpackagedomain.NewServicePackage(
		svcId,
		serviceId,
		dto.Name,
		dto.Description,
		dto.ComboDays,
		dto.Discount,
		dto.TimeInterval,
		svcpackagedomain.SvcPackageStatusAvailable,
		nil,
	)

	if err := h.cmdRepo.CreatePackage(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create service-package").
			WithInner(err.Error())
	}
	return nil
}
