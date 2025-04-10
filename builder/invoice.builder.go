package builder

import (
	"github.com/jmoiron/sqlx"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	invoicerepository "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/infars/repository"
	invoicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/commands"
	invoicequeries "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/usecase/queries"
)

type builderOfInvoice struct {
	db    *sqlx.DB
	payOS common.PayOSConfig
}

func NewInvoiceBuilder(db *sqlx.DB) builderOfInvoice {
	return builderOfInvoice{db: db}
}

func (s builderOfInvoice) AddPayOsConfig(payOS common.PayOSConfig) builderOfInvoice {
	s.payOS = payOS
	return s
}

func (s builderOfInvoice) BuildInvoiceCmdRepo() invoicecommands.InvoiceCommandRepo {
	return invoicerepository.NewInvoiceRepo(s.db)
}

func (s builderOfInvoice) BuilderPayosConfig() common.PayOSConfig {
	return s.payOS
}

func (s builderOfInvoice) BuildInvoiceQueryRepo() invoicequeries.InvoiceQueryRepo {
	return invoicerepository.NewInvoiceRepo(s.db)
}
