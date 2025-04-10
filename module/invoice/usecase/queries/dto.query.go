package invoicequeries

import (
	"time"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type InvoiceDTO struct {
	Id            uuid.UUID  `json:"id"`
	CusPackageId  uuid.UUID  `json:"cuspackage-id"`
	OrderCode     int64      `json:"order-code"`
	TotalFee      float64    `json:"total-fee"`
	PaymentStatus string     `json:"status"`
	Note          string     `json:"note"`
	PayosUrl      string     `json:"payos-url"`
	CreatedAt     *time.Time `json:"created-at"`
}

func toInvoiceDTO(data *invoicedomain.Invoice) *InvoiceDTO {
	dto := &InvoiceDTO{
		Id:            data.GetID(),
		CusPackageId:  data.GetCusPackageID(),
		OrderCode:     data.GetOrderCode(),
		TotalFee:      data.GetTotalFee(),
		PaymentStatus: data.GetPaymentStatus().String(),
		Note:          data.GetNote(),
		PayosUrl:      data.GetPayosUrl(),
		CreatedAt:     data.GetCreatedAt(),
	}
	return dto
}
