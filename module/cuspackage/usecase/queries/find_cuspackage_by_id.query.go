package cuspackagequeries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type findCusPackageByIdHandler struct {
	queryRepo CusPackageQueryRepo
}

func NewFindCusPackageByIdHandler(queryRepo CusPackageQueryRepo) *findCusPackageByIdHandler {
	return &findCusPackageByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findCusPackageByIdHandler) Handle(ctx context.Context, cuspackageId uuid.UUID) (*CusPackageDTO, error) {
	entity, err := h.queryRepo.FindCusPackage(ctx, cuspackageId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason(fmt.Sprintf("cannot found customized-package with id: %v", cuspackageId.String())).
			WithInner(err.Error())
	}

	return toCusPackageDTO(entity), nil
}
