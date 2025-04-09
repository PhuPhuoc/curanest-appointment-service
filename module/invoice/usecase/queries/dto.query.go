package invoicequeries

import (
	"time"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type InvoiceDTO struct {
	Id            uuid.UUID  `json:"id"`
	CusPackageId  uuid.UUID  `json:"cuspackage-id"`
	TotalFee      float64    `json:"total-fee"`
	PaymentStatus string     `json:"status"`
	Note          string     `json:"note"`
	CreatedAt     *time.Time `json:"created-at"`
}

func toInvoiceDTO(data *invoicedomain.Invoice) *InvoiceDTO {
	dto := &InvoiceDTO{
		Id:            data.GetID(),
		CusPackageId:  data.GetCusPackageID(),
		TotalFee:      data.GetTotalFee(),
		PaymentStatus: data.GetPaymentStatus().String(),
		Note:          data.GetNote(),
		CreatedAt:     data.GetCreatedAt(),
	}
	return dto
}
