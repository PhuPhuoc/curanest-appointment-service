package servicecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/service/domain"
)

type updateServiceHandler struct {
	cmdRepo ServiceCommandRepo
}

func NewUpdateServiceHandler(cmdRepo ServiceCommandRepo) *updateServiceHandler {
	return &updateServiceHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateServiceHandler) Handle(ctx context.Context, dto UpdateServiceDTO) error {
	newService, _ := servicedomain.NewService(
		dto.Id,
		dto.CategoryId,
		dto.Name,
		dto.Description,
		dto.EstDuration,
		servicedomain.Enum(dto.Status),
		dto.CreatedAt,
	)

	if err := h.cmdRepo.Update(ctx, newService); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update service").
			WithInner(err.Error())
	}
	return nil
}
