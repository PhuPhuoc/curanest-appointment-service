package categorycommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categorydomain "github.com/PhuPhuoc/curanest-appointment-service/module/category/domain"
)

type createCategoryHandler struct {
	cmdRepo CategoryCommandRepo
}

func NewCreateCategoryHandler(cmdRepo CategoryCommandRepo) *createCategoryHandler {
	return &createCategoryHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createCategoryHandler) Handle(ctx context.Context, dto *CreateCategoryDTO) error {
	cateId := common.GenUUID()
	entity, _ := categorydomain.NewCategory(
		cateId,
		nil,
		dto.Name,
		dto.Description,
		nil,
	)
	if err := h.cmdRepo.Create(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create category").
			WithInner(err.Error())
	}
	return nil
}
