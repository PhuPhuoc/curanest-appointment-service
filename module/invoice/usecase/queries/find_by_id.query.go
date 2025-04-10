package invoicequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type getInvoiceHandler struct {
	queryRepo InvoiceQueryRepo
}

func NewGetInvoiceHandler(queryRepo InvoiceQueryRepo) *getInvoiceHandler {
	return &getInvoiceHandler{
		queryRepo: queryRepo,
	}
}

func (h *getInvoiceHandler) Handle(ctx context.Context, Id uuid.UUID) (*InvoiceDTO, error) {
	entity, err := h.queryRepo.FindById(ctx, Id)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get invoice with id(" + Id.String() + ")").
			WithInner(err.Error())
	}

	return toInvoiceDTO(entity), nil
}
