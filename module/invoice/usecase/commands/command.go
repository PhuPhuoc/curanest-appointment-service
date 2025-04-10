package invoicecommands

import "github.com/PhuPhuoc/curanest-appointment-service/common"

type Commands struct {
	GetUrlPayment *getUrlPaymentHandler
}

type Builder interface {
	BuildInvoiceCmdRepo() InvoiceCommandRepo
	BuilderPayosConfig() common.PayOSConfig
}

func NewInvoiceCmdWithBuilder(b Builder) Commands {
	return Commands{
		GetUrlPayment: NewGetUrlPaymentHandler(
			b.BuildInvoiceCmdRepo(),
			b.BuilderPayosConfig(),
		),
	}
}

type InvoiceCommandRepo interface{}
