package cuspackagequeries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type findCustaskByIdHandler struct {
	queryRepo CusPackageQueryRepo
}

func NewFindCustaskByIdHandler(queryRepo CusPackageQueryRepo) *findCustaskByIdHandler {
	return &findCustaskByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findCustaskByIdHandler) Handle(ctx context.Context, cuspackageId uuid.UUID) (*CusTaskDTO, error) {
	entity, err := h.queryRepo.FindCusTaskById(ctx, cuspackageId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason(fmt.Sprintf("cannot found custask with id: %v", cuspackageId.String())).
			WithInner(err.Error())
	}

	return toCusTaskDTO(entity), nil
}
