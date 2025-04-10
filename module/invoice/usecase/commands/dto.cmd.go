package invoicecommands

import "github.com/google/uuid"

type DetailInvoiceCmdDTO struct {
	Id        uuid.UUID `json:"id"`
	OrderCode int64     `json:"order-code"`
	TotalFee  float64   `json:"total-fee"`
}

type UrlPayment struct {
	Url string `json:"url"`
}
