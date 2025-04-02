package invoicerepository

import (
	"context"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

func (repo *invoiceRepo) CreateInvoice(ctx context.Context, entity *invoicedomain.Invoice) error {
	return nil
}
