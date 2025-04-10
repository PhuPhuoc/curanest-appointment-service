package invoicerepository

import (
	"time"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

var (
	TABLE_INVOICE = `invoices`

	CREATE_INVOICE = []string{
		"id",
		"customized_package_id",
		"order_code",
		"total_fee",
		"payment_status",
		"payos_url",
		"note",
	}

	GET_INVOICE = []string{
		"id",
		"customized_package_id",
		"order_code",
		"total_fee",
		"payment_status",
		"note",
		"payos_url",
		"created_at",
	}

	UPDATE_INVOICE = []string{
		"payment_status",
		"note",
	}
)

type InvoiceDTO struct {
	Id            uuid.UUID  `db:"id"`
	CusPackageId  uuid.UUID  `db:"customized_package_id"`
	OrderCode     int64      `db:"order_code"`
	TotalFee      float64    `db:"total_fee"`
	PaymentStatus string     `db:"payment_status"`
	Note          string     `db:"note"`
	PayosUrl      string     `db:"payos_url"`
	CreatedAt     *time.Time `db:"created_at"`
}

func (dto *InvoiceDTO) ToInvoiceEntity() (*invoicedomain.Invoice, error) {
	return invoicedomain.NewInvoice(
		dto.Id,
		dto.CusPackageId,
		dto.OrderCode,
		dto.TotalFee,
		invoicedomain.EnumPaymentStatus(dto.PaymentStatus),
		dto.Note,
		dto.PayosUrl,
		dto.CreatedAt,
	)
}

func ToInvoiceDTO(data *invoicedomain.Invoice) *InvoiceDTO {
	return &InvoiceDTO{
		Id:            data.GetID(),
		CusPackageId:  data.GetCusPackageID(),
		OrderCode:     data.GetOrderCode(),
		TotalFee:      data.GetTotalFee(),
		PaymentStatus: data.GetPaymentStatus().String(),
		Note:          data.GetNote(),
		PayosUrl:      data.GetPayosUrl(),
	}
}
