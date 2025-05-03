package invoicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	"github.com/google/uuid"
)

func (repo *invoiceRepo) FindById(ctx context.Context, Id uuid.UUID) (*invoicedomain.Invoice, error) {
	where := "id = ?"
	query := common.GenerateSQLQueries(common.FIND, TABLE_INVOICE, GET_INVOICE, &where)
	var dto InvoiceDTO
	if err := repo.db.GetContext(ctx, &dto, query, Id); err != nil {
		return nil, err
	}

	return dto.ToInvoiceEntity()
}
