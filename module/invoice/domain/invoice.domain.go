package invoicedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	id            uuid.UUID
	cusPackageId  uuid.UUID
	totalFee      float64
	paymentStatus PaymentStatus
	paymentType   PaymentType
	note          string
	createdAt     *time.Time
}

func (a *Invoice) GetID() uuid.UUID {
	return a.id
}

func (a *Invoice) GetCusPackageID() uuid.UUID {
	return a.cusPackageId
}

func (a *Invoice) GetTotalFee() float64 {
	return a.totalFee
}

func (a *Invoice) GetPaymentStatus() PaymentStatus {
	return a.paymentStatus
}

func (a *Invoice) GetPaymentType() PaymentType {
	return a.paymentType
}

func (a *Invoice) GetNote() string {
	return a.note
}

func (a *Invoice) GetCreatedAt() *time.Time {
	return a.createdAt
}

func NewInvoice(
	id, cusPackageId uuid.UUID,
	totalFee float64,
	paymentStatus PaymentStatus,
	paymentType PaymentType,
	note string,
	createdAt *time.Time,
) (*Invoice, error) {
	return &Invoice{
		id:            id,
		cusPackageId:  cusPackageId,
		totalFee:      totalFee,
		paymentStatus: paymentStatus,
		paymentType:   paymentType,
		note:          note,
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

type PaymentType int

const (
	PaymentTypeCashToNurse PaymentType = iota
	PaymentTypeWallet
)

func EnumPaymentType(s string) PaymentType {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "wallet":
		return PaymentTypeWallet
	default:
		return PaymentTypeCashToNurse
	}
}

func (r PaymentType) String() string {
	switch r {
	case PaymentTypeWallet:
		return "wallet"
	case PaymentTypeCashToNurse:
		return "cash_to_nurse"
	default:
		return "unknown"
	}
}
