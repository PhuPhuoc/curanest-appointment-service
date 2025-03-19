package servicecommands

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

type createServiceHandler struct {
	cmdRepo ServiceCommandRepo
}

func NewCreateServiceHandler(cmdRepo ServiceCommandRepo) *createServiceHandler {
	return &createServiceHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createServiceHandler) Handle(ctx context.Context, cateId uuid.UUID, dto *CreateServiceDTO) error {
	serviceId := common.GenUUID()

	entity, _ := servicedomain.NewService(
		serviceId,
		cateId,
		dto.Name,
		dto.Description,
		dto.EstDuration,
		servicedomain.StatusAvailable,
		nil,
	)

	if err := h.cmdRepo.Create(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create service").
			WithInner(err.Error())
	}
	return nil
}
