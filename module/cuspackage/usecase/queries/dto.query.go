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
	Id            uuid.UUID  `json:"id"`
	SvcPackageId  uuid.UUID  `json:"svc-package-id"`
	PatientId     uuid.UUID  `json:"patient-id"`
	Name          string     `json:"name"`
	TotalFee      float64    `json:"total-fee"`
	PaidAmount    float64    `json:"paid-amount"`
	UnpaidAmount  float64    `json:"unpaid-amount"`
	PaymentStatus string     `json:"payment-status"`
	CreatedAt     *time.Time `json:"created-at"`
}

func toCusPackageDTO(data *cuspackagedomain.CustomizedPackage) *CusPackageDTO {
	dto := &CusPackageDTO{
		Id:            data.GetID(),
		Name:          data.GetName(),
		SvcPackageId:  data.GetServicePackageID(),
		PatientId:     data.GetPatientID(),
		TotalFee:      data.GetTotalFee(),
		PaidAmount:    data.GetPaidAmount(),
		UnpaidAmount:  data.GetUnpaidAmount(),
		PaymentStatus: data.GetPaymentStatus().String(),
		CreatedAt:     data.GetCreatedAt(),
	}
	return dto
}

type CusTaskDTO struct {
	Id           uuid.UUID `json:"id"`
	SvcTaskId    uuid.UUID `json:"-"`
	CusPackageId uuid.UUID `json:"-"`
	TaskOrder    int       `json:"task-order"`
	Name         string    `json:"name"`
	ClientNote   string    `json:"client-note"`
	StaffAdvice  string    `json:"staff-advice"`
	EstDuration  int       `json:"est-duration"`
	Unit         string    `json:"unit"`
	TotalUnit    int       `json:"total-unit"`
	TotalCost    float64   `json:"total-cost"`
	Status       string    `json:"status"`
	EstDate      time.Time `json:"est-date"`
}

func (ct *CusTaskDTO) ToCusTaskEntity() (*cuspackagedomain.CustomizedTask, error) {
	return cuspackagedomain.NewCustomizedTask(
		ct.Id,
		ct.SvcTaskId,
		ct.CusPackageId,
		ct.TaskOrder,
		ct.Name,
		ct.ClientNote,
		ct.StaffAdvice,
		ct.EstDuration,
		ct.TotalCost,
		cuspackagedomain.EnumCusTaskUnit(ct.Unit),
		ct.TotalUnit,
		ct.EstDate,
		nil,
		cuspackagedomain.EnumCusTaskStatus(ct.Status),
	)
}

func toCusTaskDTO(data *cuspackagedomain.CustomizedTask) *CusTaskDTO {
	dto := &CusTaskDTO{
		Id:           data.GetID(),
		SvcTaskId:    data.GetSvcTaskID(),
		CusPackageId: data.GetCusPackageID(),
		TaskOrder:    data.GetTaskOrder(),
		Name:         data.GetName(),
		ClientNote:   data.GetClientNote(),
		StaffAdvice:  data.GetStaffAdvice(),
		EstDuration:  data.GetEstDuration(),
		Unit:         data.GetUnit().String(),
		TotalUnit:    data.GetTotalUnit(),
		Status:       data.GetStatus().String(),
		EstDate:      data.GetEstDate(),
	}
	return dto
}

type MedicalRecordDTO struct {
	Id                uuid.UUID  `json:"id"`
	NursingId         *uuid.UUID `json:"svc-package-id"`
	AppointmentId     uuid.UUID  `json:"patient-id"`
	NursingReport     string     `json:"nursing-report"`
	StaffConfirmation string     `json:"staff-confirmation"`
	Status            string     `json:"status"`
	CreatedAt         *time.Time `json:"created-at"`
}

func (m *MedicalRecordDTO) ToMedicalRecordEntity() (*cuspackagedomain.MedicalRecord, error) {
	return cuspackagedomain.NewMedicalRecord(
		m.Id,
		m.AppointmentId,
		m.NursingId,
		m.NursingReport,
		m.StaffConfirmation,
		cuspackagedomain.EnumRecordStatus(m.Status),
		m.CreatedAt,
	)
}

func toMedicalRecordDTO(data *cuspackagedomain.MedicalRecord) *MedicalRecordDTO {
	dto := &MedicalRecordDTO{
		Id:                data.GetID(),
		NursingId:         data.GetNursingId(),
		AppointmentId:     data.GetAppointmentId(),
		NursingReport:     data.GetNursingReport(),
		StaffConfirmation: data.GetStaffConfirm(),
		Status:            data.GetStatus().String(),
		CreatedAt:         data.GetCreatedAt(),
	}
	return dto
}
