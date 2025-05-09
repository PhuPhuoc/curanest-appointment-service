package invoicecommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
	"github.com/google/uuid"
)

type Commands struct {
	WebHookGoong     *webhookGoongHandler
	CancelPaymentUrl *cancelPaymentUrlHandler
	CreateNewUrl     *createNewPaymentUrlHandler
}

type Builder interface {
	BuildInvoiceCmdRepo() InvoiceCommandRepo
	BuildCusPackageFetcher() CusPackageFetcher
	BuilderPayosConfig() common.PayOSConfig
}

func NewInvoiceCmdWithBuilder(b Builder) Commands {
	return Commands{
		WebHookGoong: NewWebhoobGoongHandler(
			b.BuildInvoiceCmdRepo(),
			b.BuildCusPackageFetcher(),
		),
		CancelPaymentUrl: NewCancelPaymentUrlHandler(
			b.BuildInvoiceCmdRepo(),
		),
		CreateNewUrl: NewCreateNewPaymentUrlHandler(
			b.BuildInvoiceCmdRepo(),
			b.BuilderPayosConfig(),
		),
	}
}

type InvoiceCommandRepo interface {
	UpdateInvoiceFromPayos(ctx context.Context, orderCode string) error
	UpdateInvoice(ctx context.Context, entity *invoicedomain.Invoice) error
}

type CusPackageFetcher interface {
	FindCusPackage(ctx context.Context, id uuid.UUID) (*cuspackagedomain.CustomizedPackage, error)
	UpdateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
}
