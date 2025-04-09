package cuspackagequeries

import (
	"time"

	"github.com/google/uuid"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

type FilterGetCusPackageTaskDTO struct {
	CusPackageId uuid.UUID `json:"cus-package-id"`
	EstDate      time.Time `json:"est-date"`
}

type PackageTaskResponse struct {
	Package *CusPackageDTO `json:"package"`
	Tasks   []CusTaskDTO   `json:"tasks"`
}

type CusPackageDTO struct {
	Id uuid.UUID `json:"id"`
	// SvcPackageId uuid.UUID `json:"svc-package-id"`
	// PatientId    uuid.UUID `json:"patient-id"`
	Name          string    `json:"name"`
	TotalFee      float64   `json:"total-fee"`
	PaidAmount    float64   `json:"paid-amount"`
	UnpaidAmount  float64   `json:"unpaid-amount"`
	PaymentStatus string    `json:"payment-status"`
	CreatedAt     time.Time `json:"created-at"`
}

func toCusPackageDTO(data *cuspackagedomain.CustomizedPackage) *CusPackageDTO {
	dto := &CusPackageDTO{
		Id:            data.GetID(),
		Name:          data.GetName(),
		TotalFee:      data.GetTotalFee(),
		PaidAmount:    data.GetPaidAmount(),
		UnpaidAmount:  data.GetUnpaidAmount(),
		PaymentStatus: data.GetPaymentStatus().String(),
		CreatedAt:     data.GetCreatedAt(),
	}
	return dto
}

type CusTaskDTO struct {
	Id          uuid.UUID `json:"id"`
	TaskOrder   int       `json:"task-order"`
	Name        string    `json:"name"`
	ClientNote  string    `json:"client-note"`
	StaffAdvice string    `json:"staff-advice"`
	EstDuration int       `json:"est-duration"`
	Unit        string    `json:"unit"`
	TotalUnit   int       `json:"total-unit"`
	Status      string    `json:"status"`
	EstDate     time.Time `json:"est-date"`
}

func toCusTaskDTO(data *cuspackagedomain.CustomizedTask) *CusTaskDTO {
	dto := &CusTaskDTO{
		Id:          data.GetID(),
		TaskOrder:   data.GetTaskOrder(),
		Name:        data.GetName(),
		ClientNote:  data.GetClientNote(),
		StaffAdvice: data.GetStaffAdvice(),
		EstDuration: data.GetEstDuration(),
		Unit:        data.GetUnit().String(),
		TotalUnit:   data.GetTotalUnit(),
		Status:      data.GetStatus().String(),
		EstDate:     data.GetEstDate(),
	}
	return dto
}
