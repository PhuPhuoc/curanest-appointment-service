package invoicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

func (repo *invoiceRepo) UpdateInvoice(ctx context.Context, entity *invoicedomain.Invoice) error {
	dto := ToInvoiceDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_INVOICE, UPDATE_INVOICE, &where)

	// If no transaction, use db directly
	_, err := repo.db.NamedExec(query, dto)
	return err
}
