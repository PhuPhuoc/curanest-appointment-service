package svcpackagecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type updateTaskOrderHandler struct {
	cmdRepo SvcPackageCommandRepo
}

func NewUpdateTaskOrderHandler(cmdRepo SvcPackageCommandRepo) *updateTaskOrderHandler {
	return &updateTaskOrderHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateTaskOrderHandler) Handle(ctx context.Context, dtos *UpdateTaskOrderDTO) error {
	entities := make([]svcpackagedomain.ServiceTask, len(dtos.SvcTasks))
	for i := range dtos.SvcTasks {
		entity, _ := svcpackagedomain.NewServiceTask(
			dtos.SvcTasks[i].Id,
			dtos.SvcTasks[i].SvcPackageId,
			dtos.SvcTasks[i].IsMustHave,
			0,
			dtos.SvcTasks[i].Name,
			dtos.SvcTasks[i].Description,
			dtos.SvcTasks[i].StaffAdvice,
			dtos.SvcTasks[i].EstDuration,
			dtos.SvcTasks[i].Cost,
			dtos.SvcTasks[i].AdditionalCost,
			dtos.SvcTasks[i].AdditionalCostDesc,
			svcpackagedomain.EnumSvcTaskUnit(dtos.SvcTasks[i].Unit),
			dtos.SvcTasks[i].PriceOfStep,
			svcpackagedomain.EnumSvcTaskStatus(dtos.SvcTasks[i].Status),
		)
		entities[i] = *entity
	}

	if err := h.cmdRepo.UpdateTaskOrder(ctx, entities); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update task-order of list service-task").
			WithInner(err.Error())
	}
	return nil
}
