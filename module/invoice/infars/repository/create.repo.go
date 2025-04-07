package invoicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

func (repo *invoiceRepo) CreateInvoice(ctx context.Context, entity *invoicedomain.Invoice) error {
	dto := ToInvoiceDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE_INVOICE, CREATE_INVOICE, nil)

	// Get transaction from context if exist
	if tx := common.GetTxFromContext(ctx); tx != nil {
		_, err := tx.NamedExec(query, dto)
		return err
	}

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}
