package categoryqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type findCategoryByIdHandler struct {
	queryRepo CategoryQueryRepo
}

func NewFindCategoryByIdHandler(queryRepo CategoryQueryRepo) *findCategoryByIdHandler {
	return &findCategoryByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findCategoryByIdHandler) Handle(ctx context.Context, cateId uuid.UUID) (*CategoryDTO, error) {
	entity, err := h.queryRepo.FindCategoryById(ctx, cateId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get this category info").
			WithInner(err.Error())
	}

	return toDTO(entity), nil
}
