package invoicecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type cancelPaymentUrlHandler struct {
	cmdRepo InvoiceCommandRepo
}

func NewCancelPaymentUrlHandler(cmdRepo InvoiceCommandRepo) *cancelPaymentUrlHandler {
	return &cancelPaymentUrlHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *cancelPaymentUrlHandler) Handle(ctx context.Context, e *invoicedomain.Invoice) error {
	if e.GetPaymentStatus() == invoicedomain.PaymentStatusPaid {
		return common.NewInternalServerError().
			WithReason("cannot cannel payment url because this invoice has beed paid")
	}

	newInvoice, _ := invoicedomain.NewInvoice(
		e.GetID(),
		e.GetCusPackageID(),
		nil,
		e.GetTotalFee(),
		e.GetPaymentStatus(),
		e.GetNote(),
		nil,
		e.GetCreatedAt(),
	)

	if err := h.cmdRepo.UpdateInvoice(ctx, newInvoice); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update this invoice").
			WithInner(err.Error())
	}

	return nil
}
