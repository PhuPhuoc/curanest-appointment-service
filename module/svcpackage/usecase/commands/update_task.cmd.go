package svcpackagecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type updateSvcTaskHandler struct {
	cmdRepo SvcPackageCommandRepo
}

func NewUpdateSvcTaskHandler(cmdRepo SvcPackageCommandRepo) *updateSvcTaskHandler {
	return &updateSvcTaskHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateSvcTaskHandler) Handle(ctx context.Context, dto *UpdateServiceTaskDTO) error {
	entity, _ := svcpackagedomain.NewServiceTask(
		dto.SvcTaskId,
		dto.SvcPackageId,
		dto.IsMustHave,
		dto.Order,
		dto.Name,
		dto.Description,
		dto.StaffAdvice,
		dto.EstDuration,
		dto.Cost,
		dto.AdditionalCost,
		dto.AdditionalCostDesc,
		svcpackagedomain.EnumSvcTaskUnit(dto.Unit),
		dto.PriceOfStep,
		svcpackagedomain.EnumSvcTaskStatus(dto.Status),
	)

	if err := h.cmdRepo.UpdateTask(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update service-task").
			WithInner(err.Error())
	}
	return nil
}
