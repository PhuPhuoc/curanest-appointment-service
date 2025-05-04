package invoicequeries

import (
	"context"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type Queries struct {
	FindInvoice            *findInvoiceHandler
	GetInvoiceById         *getInvoiceHandler
	GetInvoiceByOrderCode  *getInvoiceByOrderCodeHandler
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
		GetInvoiceByOrderCode: NewGetInvoiceByOrderCodeHandler(
			b.BuildInvoiceQueryRepo(),
		),
		GetInvoiceByPatientIds: NewGetInvoicesByPatientIdHandler(
			b.BuildInvoiceQueryRepo(),
		),
	}
}

type InvoiceQueryRepo interface {
	FindById(ctx context.Context, Id uuid.UUID) (*invoicedomain.Invoice, error)
	FindByOrderCode(ctx context.Context, ordercode int64) (*invoicedomain.Invoice, error)
	FindByCusPackageId(ctx context.Context, cusPackageId uuid.UUID) ([]invoicedomain.Invoice, error)
	GetInvoicesByPatientId(ctx context.Context, isAdmin bool, patientIds []uuid.UUID) ([]invoicedomain.Invoice, error)
}
