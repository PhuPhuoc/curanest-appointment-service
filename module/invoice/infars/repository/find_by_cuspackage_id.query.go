package invoicerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	"github.com/google/uuid"
)

func (repo *invoiceRepo) FindByCusPackageId(ctx context.Context, cusPackageId uuid.UUID) ([]invoicedomain.Invoice, error) {
	where := "customized_package_id = ?"
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_INVOICE, GET_INVOICE, &where)
	var dtos []InvoiceDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, cusPackageId); err != nil {
		return nil, err
	}

	entities := make([]invoicedomain.Invoice, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToInvoiceEntity()
		entities[i] = *entity
	}

	return entities, nil
}
