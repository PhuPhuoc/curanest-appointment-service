package svcpackagecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type updateSvcPackageHandler struct {
	cmdRepo SvcPackageCommandRepo
}

func NewUpdateSvcPackageHandler(cmdRepo SvcPackageCommandRepo) *updateSvcPackageHandler {
	return &updateSvcPackageHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateSvcPackageHandler) Handle(ctx context.Context, dto *UpdateServicePackageDTO) error {
	entity, _ := svcpackagedomain.NewServicePackage(
		dto.SvcPackageId,
		dto.ServiceId,
		dto.Name,
		dto.Description,
		dto.ComboDays,
		dto.Discount,
		dto.TimeInterval,
		svcpackagedomain.EnumSvcPackageStatus(dto.Status),
		nil,
	)

	if err := h.cmdRepo.UpdatePackage(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update service-package").
			WithInner(err.Error())
	}
	return nil
}
