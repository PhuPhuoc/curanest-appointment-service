package invoicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

func (repo *invoiceRepo) FindByOrderCode(ctx context.Context, ordercode int64) (*invoicedomain.Invoice, error) {
	where := "order_code = ?"
	query := common.GenerateSQLQueries(common.FIND, TABLE_INVOICE, GET_INVOICE, &where)
	var dto InvoiceDTO
	if err := repo.db.GetContext(ctx, &dto, query, ordercode); err != nil {
		return nil, err
	}

	return dto.ToInvoiceEntity()
}
