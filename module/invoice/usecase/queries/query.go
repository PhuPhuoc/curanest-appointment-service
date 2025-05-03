package invoicequeries

import (
	"context"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type Queries struct {
	FindInvoice            *findInvoiceHandler
	GetInvoiceById         *getInvoiceHandler
	GetInvoiceByPatientIds *getInvoicesByPatientIdHandler
}

type Builder interface {
	BuildInvoiceQueryRepo() InvoiceQueryRepo
}

func NewInvoiceQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindInvoice: NewFindInvoiceHandler(
			b.BuildInvoiceQueryRepo(),
		),
		GetInvoiceById: NewGetInvoiceHandler(
			b.BuildInvoiceQueryRepo(),
		),
		GetInvoiceByPatientIds: NewGetInvoicesByPatientIdHandler(
			b.BuildInvoiceQueryRepo(),
		),
	}
}

type InvoiceQueryRepo interface {
	FindById(ctx context.Context, Id uuid.UUID) (*invoicedomain.Invoice, error)
	FindByCusPackageId(ctx context.Context, cusPackageId uuid.UUID) ([]invoicedomain.Invoice, error)
	GetInvoicesByPatientId(ctx context.Context, isAdmin bool, patientIds []uuid.UUID) ([]invoicedomain.Invoice, error)
}
