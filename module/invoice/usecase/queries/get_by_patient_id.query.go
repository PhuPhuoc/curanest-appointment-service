package invoicequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getInvoicesByPatientIdHandler struct {
	queryRepo InvoiceQueryRepo
}

func NewGetInvoicesByPatientIdHandler(queryRepo InvoiceQueryRepo) *getInvoicesByPatientIdHandler {
	return &getInvoicesByPatientIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getInvoicesByPatientIdHandler) Handle(ctx context.Context, req RequestGetInvoicesByPatientIds) ([]InvoiceDTO, error) {
	entities, err := h.queryRepo.GetInvoicesByPatientId(ctx, req.IsAdmin, req.PatientIds)
	if err != nil {
		return []InvoiceDTO{}, common.NewInternalServerError().
			WithReason("cannot get invoices with patientids").
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
