package invoicequeries

import (
	"time"

	"github.com/google/uuid"

	invoicedomain "github.com/PhuPhuoc/curanest-appointment-service/module/invoice/domain"
)

type RequestGetInvoicesByPatientIds struct {
	PatientIds []uuid.UUID `json:"patient-ids"`
	IsAdmin    bool        `json:"is-admin"`
}

type InvoiceDTO struct {
	Id            uuid.UUID  `json:"id"`
	CusPackageId  uuid.UUID  `json:"cuspackage-id"`
	OrderCode     int64      `json:"order-code,omitempty"`
	TotalFee      float64    `json:"total-fee"`
	PaymentStatus string     `json:"status"`
	Note          string     `json:"note,omitempty"`
	PayosUrl      string     `json:"payos-url,omitempty"`
	CreatedAt     *time.Time `json:"created-at"`
}

func (i *InvoiceDTO) ToInvoiceEntity() (*invoicedomain.Invoice, error) {
	return invoicedomain.NewInvoice(
		i.Id,
		i.CusPackageId,
		i.OrderCode,
		i.TotalFee,
		invoicedomain.EnumPaymentStatus(i.PaymentStatus),
		i.Note,
		i.PayosUrl,
		i.CreatedAt,
	)
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
