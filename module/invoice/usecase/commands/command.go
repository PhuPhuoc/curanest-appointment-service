package invoicecommands

type Commands struct{}

type Builder interface {
	BuildInvoiceCmdRepo() InvoiceCommandRepo
}

func NewInvoiceCmdWithBuilder(b Builder) Commands {
	return Commands{}
}

type InvoiceCommandRepo interface{}
