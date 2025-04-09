package invoicequeries

import (
	"context"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type Queries struct {
	FindInvoice *findInvoiceHandler
}

type Builder interface {
	BuildInvoiceQueryRepo() InvoiceQueryRepo
}

func NewInvoiceQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindInvoice: NewFindInvoiceHandler(
			b.BuildInvoiceQueryRepo(),
		),
	}
}

type InvoiceQueryRepo interface {
	FindByCusPackageId(ctx context.Context, cusPackageId uuid.UUID) ([]invoicedomain.Invoice, error)
}
