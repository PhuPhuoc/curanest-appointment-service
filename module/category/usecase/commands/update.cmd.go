package categorycommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type updateCategoryHandler struct {
	cmdRepo CategoryCommandRepo
}

func NewUpdateCategoryHandler(cmdRepo CategoryCommandRepo) *updateCategoryHandler {
	return &updateCategoryHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updateCategoryHandler) Handle(ctx context.Context, dto *UpdateCategoryDTO) error {
	entity, _ := categorydomain.NewCategory(
		dto.Id,
		dto.StaffId,
		dto.Name,
		dto.Description,
		dto.Thumbnail,
		dto.CreatedAt,
	)
	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create category").
			WithInner(err.Error())
	}
	return nil
}
