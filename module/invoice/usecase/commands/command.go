package invoicecommands

import "context"

type Commands struct {
	WebHookGoong *webhookGoongHandler
}

type Builder interface {
	BuildInvoiceCmdRepo() InvoiceCommandRepo
}

func NewInvoiceCmdWithBuilder(b Builder) Commands {
	return Commands{
		WebHookGoong: NewWebhoobGoongHandler(
			b.BuildInvoiceCmdRepo(),
		),
	}
}

type InvoiceCommandRepo interface {
	UpdateInvoiceFromGoong(ctx context.Context, orderCode string) error
}
