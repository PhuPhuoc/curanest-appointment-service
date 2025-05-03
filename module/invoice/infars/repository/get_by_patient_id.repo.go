package invoicerepository

import (
	"context"
	"fmt"
	"strings"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	"github.com/google/uuid"
)

func (repo *invoiceRepo) GetInvoicesByPatientId(ctx context.Context, patientIds []uuid.UUID) ([]invoicedomain.Invoice, error) {
	if len(patientIds) == 0 {
		return []invoicedomain.Invoice{}, fmt.Errorf("patientIds cannot be empty")
	}

	patientIdStrs := make([]string, len(patientIds))
	for i, id := range patientIds {
		patientIdStrs[i] = id.String()
	}
	patientIdsParam := "'" + strings.Join(patientIdStrs, "','") + "'"

	query := `
		select i.id, i.customized_package_id, i.total_fee, i.payment_status, i.created_at from invoices i
		join customized_packages cp on i.customized_package_id = cp.id
		where cp.patient_id in (%s)
	`
	query = fmt.Sprintf(query, patientIdsParam)

	dtos := []InvoiceDTO{}
	if err := repo.db.SelectContext(ctx, &dtos, query); err != nil {
		return nil, fmt.Errorf("failed to get list invoices: %w", err)
	}

	entities := make([]invoicedomain.Invoice, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToInvoiceEntity()
		entities[i] = *entity
	}
	return entities, nil
}
