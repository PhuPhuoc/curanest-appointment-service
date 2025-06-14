package cuspackagedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type CustomizedPackage struct {
	id            uuid.UUID
	svcPackageId  uuid.UUID
	patientId     uuid.UUID
	name          string
	totalFee      float64
	paidAmount    float64
	uppaidAmount  float64
	paymentStatus PaymentStatus
	isCancel      bool
	createdAt     *time.Time
}

func (a *CustomizedPackage) GetID() uuid.UUID {
	return a.id
}

func (a *CustomizedPackage) GetServicePackageID() uuid.UUID {
	return a.svcPackageId
}

func (a *CustomizedPackage) GetPatientID() uuid.UUID {
	return a.patientId
}

func (a *CustomizedPackage) GetName() string {
	return a.name
}

func (a *CustomizedPackage) GetTotalFee() float64 {
	return a.totalFee
}

func (a *CustomizedPackage) GetPaidAmount() float64 {
	return a.paidAmount
}

func (a *CustomizedPackage) GetUnpaidAmount() float64 {
	return a.uppaidAmount
}

func (a *CustomizedPackage) GetPaymentStatus() PaymentStatus {
	return a.paymentStatus
}

func (a *CustomizedPackage) GetIsCancel() bool {
	return a.isCancel
}

func (a *CustomizedPackage) GetCreatedAt() *time.Time {
	return a.createdAt
}

func NewCustomizedPackage(
	id, svcPackageId, patientId uuid.UUID,
	name string,
	totalFee, paidAmount, unpaidAmount float64,
	paymentStatus PaymentStatus,
	isCancel bool,
	createdAt *time.Time,
) (*CustomizedPackage, error) {
	return &CustomizedPackage{
		id:            id,
		svcPackageId:  svcPackageId,
		patientId:     patientId,
		name:          name,
		totalFee:      totalFee,
		paidAmount:    paidAmount,
		uppaidAmount:  unpaidAmount,
		paymentStatus: paymentStatus,
		isCancel:      isCancel,
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
