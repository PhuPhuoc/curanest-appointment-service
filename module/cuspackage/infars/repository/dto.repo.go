package cuspackagerepository

import (
	"time"

	"github.com/google/uuid"

	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

var (
	TABLE_CUSPACKAGE = `customized_packages`
	TABLE_CUSTASK    = `customized_tasks`

	CREATE_CUSPACKAGE = []string{
		"id",
		"service_package_id",
		"patient_id",
		"name",
		"total_fee",
		"paid_amount",
		"unpaid_amount",
		"payment_status",
	}
	CREATE_CUSTASK = []string{
		"id",
		"service_task_id",
		"customized_package_id",
		"task_order",
		"name",
		"client_note",
		"staff_advice",
		"est_duration",
		"total_cost",
		"total_unit",
		"unit",
		"est_date",
		"act_date",
		"status",
	}

	GET_CUSPACKAGE = []string{
		"id",
		"service_package_id",
		"patient_id",
		"name",
		"total_fee",
		"paid_amount",
		"unpaid_amount",
		"created_at",
	}
	GET_CUSTASK = []string{
		"id",
		"service_task_id",
		"customized_package_id",
		"task_order",
		"name",
		"client_note",
		"staff_advice",
		"est_duration",
		"total_cost",
		"total_unit",
		"unit",
		"est_date",
		"act_date",
		"status",
	}

	UPDATE_CUSPACKAGE = []string{
		"total_fee",
		"paid_amount",
		"unpaid_amount",
		"payment_status",
	}
	UPDATE_FIELD_TASK = []string{
		"client_note",
		"est_duration",
		"total_cost",
		"total_unit",
		"unit",
		"act_date",
		"status",
	}
)

type CusPackageDTO struct {
	Id               uuid.UUID  `db:"id"`
	ServicePackageId uuid.UUID  `db:"service_package_id"`
	PatientId        uuid.UUID  `db:"patient_id"`
	Name             string     `db:"name"`
	TotalFee         float64    `db:"total_fee"`
	PaidAmount       float64    `db:"paid_amount"`
	UnpaidAmount     float64    `db:"unpaid_amount"`
	PaymentStatus    string     `db:"payment_status"`
	CreatedAt        *time.Time `db:"created_at"`
}

func (dto *CusPackageDTO) ToCusPackageEntity() (*cuspackagedomain.CustomizedPackage, error) {
	return cuspackagedomain.NewCustomizedPackage(
		dto.Id,
		dto.ServicePackageId,
		dto.PatientId,
		dto.Name,
		dto.TotalFee,
		dto.PaidAmount,
		dto.UnpaidAmount,
		cuspackagedomain.EnumPaymentStatus(dto.PaymentStatus),
		dto.CreatedAt,
	)
}

func ToCusPackageDTO(data *cuspackagedomain.CustomizedPackage) *CusPackageDTO {
	return &CusPackageDTO{
		Id:               data.GetID(),
		ServicePackageId: data.GetServicePackageID(),
		PatientId:        data.GetPatientID(),
		Name:             data.GetName(),
		TotalFee:         data.GetTotalFee(),
		PaidAmount:       data.GetPaidAmount(),
		UnpaidAmount:     data.GetUnpaidAmount(),
		PaymentStatus:    data.GetPaymentStatus().String(),
	}
}

type CusTaskDTO struct {
	Id           uuid.UUID  `db:"id"`
	SvcTaskId    uuid.UUID  `db:"service_task_id"`
	CusPackageId uuid.UUID  `db:"customized_package_id"`
	TaskOrder    int        `db:"task_order"`
	Name         string     `db:"name"`
	ClientNote   string     `db:"client_note"`
	StaffAdvice  string     `db:"staff_advice"`
	EstDuration  int        `db:"est_duration"`
	TotalCost    float64    `db:"total_cost"`
	TotalUnit    int        `db:"total_unit"`
	Unit         string     `db:"unit"`
	EstDate      time.Time  `db:"est_date"`
	ActDate      *time.Time `db:"act_date"`
	Status       string     `db:"status"`
}

func (dto *CusTaskDTO) ToCusTaskEntity() (*cuspackagedomain.CustomizedTask, error) {
	return cuspackagedomain.NewCustomizedTask(
		dto.Id,
		dto.SvcTaskId,
		dto.CusPackageId,
		dto.TaskOrder,
		dto.Name,
		dto.ClientNote,
		dto.StaffAdvice,
		dto.EstDuration,
		dto.TotalCost,
		cuspackagedomain.EnumCusTaskUnit(dto.Unit),
		dto.TotalUnit,
		dto.EstDate,
		dto.ActDate,
		cuspackagedomain.EnumCusTaskStatus(dto.Status),
	)
}

func ToCusTaskDTO(data *cuspackagedomain.CustomizedTask) *CusTaskDTO {
	return &CusTaskDTO{
		Id:           data.GetID(),
		SvcTaskId:    data.GetSvcTaskID(),
		CusPackageId: data.GetCusPackageID(),
		TaskOrder:    data.GetTaskOrder(),
		Name:         data.GetName(),
		ClientNote:   data.GetClientNote(),
		StaffAdvice:  data.GetStaffAdvice(),
		EstDuration:  data.GetEstDuration(),
		TotalCost:    data.GetTotalCost(),
		TotalUnit:    data.GetTotalUnit(),
		Unit:         data.GetUnit().String(),
		EstDate:      data.GetEstDate(),
		ActDate:      data.GetActDate(),
		Status:       data.GetStatus().String(),
	}
}
