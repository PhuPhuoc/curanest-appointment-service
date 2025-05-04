package invoicequeries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getInvoiceByOrderCodeHandler struct {
	queryRepo InvoiceQueryRepo
}

func NewGetInvoiceByOrderCodeHandler(queryRepo InvoiceQueryRepo) *getInvoiceByOrderCodeHandler {
	return &getInvoiceByOrderCodeHandler{
		queryRepo: queryRepo,
	}
}

func (h *getInvoiceByOrderCodeHandler) Handle(ctx context.Context, ordercode int64) (*InvoiceDTO, error) {
	entity, err := h.queryRepo.FindByOrderCode(ctx, ordercode)
	if err != nil {
		mess := fmt.Sprintf("cannot get invoice with ordercode: %v", ordercode)
		return nil, common.NewInternalServerError().
			WithReason(mess).
			WithInner(err.Error())
	}

	return toInvoiceDTO(entity), nil
}
