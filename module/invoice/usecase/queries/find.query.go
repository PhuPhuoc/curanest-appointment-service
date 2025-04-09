package invoicequeries

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type findInvoiceHandler struct {
	queryRepo InvoiceQueryRepo
}

func NewFindInvoiceHandler(queryRepo InvoiceQueryRepo) *findInvoiceHandler {
	return &findInvoiceHandler{
		queryRepo: queryRepo,
	}
}

func (h *findInvoiceHandler) Handle(ctx context.Context, cusPacakgeId uuid.UUID) ([]InvoiceDTO, error) {
	entities, err := h.queryRepo.FindByCusPackageId(ctx, cusPacakgeId)
	if err != nil {
		return []InvoiceDTO{}, common.NewInternalServerError().
			WithReason("cannot get invoice with customized-package-id(" + cusPacakgeId.String() + ")").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []InvoiceDTO{}, nil
	}

	dtos := make([]InvoiceDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = *toInvoiceDTO(&entity)
	}

	return dtos, nil
}
