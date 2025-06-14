package invoicedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	id            uuid.UUID
	cusPackageId  uuid.UUID
	orderCode     *int64
	totalFee      float64
	paymentStatus PaymentStatus
	note          string
	payosUrl      *string
	qrCode        *string
	createdAt     *time.Time
}

func (a *Invoice) GetID() uuid.UUID {
	return a.id
}

func (a *Invoice) GetCusPackageID() uuid.UUID {
	return a.cusPackageId
}

func (a *Invoice) GetOrderCode() *int64 {
	return a.orderCode
}

func (a *Invoice) GetTotalFee() float64 {
	return a.totalFee
}

func (a *Invoice) GetPaymentStatus() PaymentStatus {
	return a.paymentStatus
}

func (a *Invoice) GetNote() string {
	return a.note
}

func (a *Invoice) GetPayosUrl() *string {
	return a.payosUrl
}

func (a *Invoice) GetQrCode() *string {
	return a.qrCode
}

func (a *Invoice) GetCreatedAt() *time.Time {
	return a.createdAt
}

func NewInvoice(
	id, cusPackageId uuid.UUID,
	orderCode *int64,
	totalFee float64,
	paymentStatus PaymentStatus,
	note string,
	payosUrl, qrCode *string,
	createdAt *time.Time,
) (*Invoice, error) {
	return &Invoice{
		id:            id,
		cusPackageId:  cusPackageId,
		orderCode:     orderCode,
		totalFee:      totalFee,
		paymentStatus: paymentStatus,
		note:          note,
		payosUrl:      payosUrl,
		qrCode:        qrCode,
		createdAt:     createdAt,
	}, nil
}

type PaymentStatus int

const (
	PaymentStatusUnpaid PaymentStatus = iota
	PaymentStatusPaid
)

func EnumPaymentStatus(s string) PaymentStatus {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "paid":
		return PaymentStatusPaid
	default:
		return PaymentStatusUnpaid
	}
}

func (r PaymentStatus) String() string {
	switch r {
	case PaymentStatusUnpaid:
		return "unpaid"
	case PaymentStatusPaid:
		return "paid"
	default:
		return "unknown"
	}
}
