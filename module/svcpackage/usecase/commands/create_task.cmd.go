package svcpackagecommands

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type createSvcTaskHandler struct {
	cmdRepo SvcPackageCommandRepo
}

func NewCreateSvcTaskHandler(cmdRepo SvcPackageCommandRepo) *createSvcTaskHandler {
	return &createSvcTaskHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createSvcTaskHandler) Handle(ctx context.Context, svcPackageId uuid.UUID, dto *ServiceTaskDTO) error {
	taskId := common.GenUUID()

	entity, _ := svcpackagedomain.NewServiceTask(
		taskId,
		svcPackageId,
		dto.IsMustHave,
		dto.TaskOrder,
		dto.Name,
		dto.Description,
		dto.StaffAdvice,
		dto.EstDuration,
		dto.Cost,
		dto.AdditionalCost,
		dto.AdditionalCostDesc,
		svcpackagedomain.EnumSvcTaskUnit(dto.Unit),
		dto.PriceOfStep,
		svcpackagedomain.SvcTaskStatusAvailable,
	)

	if err := h.cmdRepo.CreateTask(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create service-task").
			WithInner(err.Error())
	}
	return nil
}
