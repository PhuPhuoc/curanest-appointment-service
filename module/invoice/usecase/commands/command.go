package invoicecommands

import (
	"context"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

type Commands struct {
	WebHookGoong *webhookGoongHandler
}

type Builder interface {
	BuildInvoiceCmdRepo() InvoiceCommandRepo
	BuildCusPackageFetcher() CusPackageFetcher
}

func NewInvoiceCmdWithBuilder(b Builder) Commands {
	return Commands{
		WebHookGoong: NewWebhoobGoongHandler(
			b.BuildInvoiceCmdRepo(),
			b.BuildCusPackageFetcher(),
		),
	}
}

type InvoiceCommandRepo interface {
	UpdateInvoiceFromGoong(ctx context.Context, orderCode string) error
}

type CusPackageFetcher interface {
	FindCusPackage(ctx context.Context, id uuid.UUID) (*cuspackagedomain.CustomizedPackage, error)
	UpdateCustomizedPackage(ctx context.Context, entity *cuspackagedomain.CustomizedPackage) error
}
