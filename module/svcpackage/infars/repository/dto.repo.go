package svcpackagerepository

import (
	"time"

	"github.com/google/uuid"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

var (
	TABLE_PACKAGE = `service_packages`
	TABLE_TASK    = `service_tasks`

	CREATE_FIELD_PACKAGE = []string{
		"id",
		"service_id",
		"name",
		"description",
		"combo_days",
		"discount",
		"time_interval",
		"status",
	}
	CREATE_FIELD_TASK = []string{
		"id",
		"service_package_id",
		"is_must_have",
		"task_order",
		"name",
		"description",
		"staff_advice",
		"est_duration",
		"cost",
		"additional_cost",
		"additional_cost_desc",
		"unit",
		"price_of_step",
		"status",
	}

	GET_FIELD_PACKAGE = []string{
		"id",
		"service_id",
		"name",
		"description",
		"combo_days",
		"discount",
		"time_interval",
		"status",
		"created_at",
	}
	GET_FIELD_TASK = []string{
		"id",
		"service_package_id",
		"is_must_have",
		"task_order",
		"name",
		"description",
		"staff_advice",
		"est_duration",
		"cost",
		"additional_cost",
		"additional_cost_desc",
		"unit",
		"price_of_step",
		"status",
	}

	UPDATE_FIELD_PACKAGE = []string{
		"name",
		"description",
		"combo_days",
		"discount",
		"time_interval",
		"status",
	}
	UPDATE_FIELD_TASK = []string{
		"is_must_have",
		"name",
		"description",
		"staff_advice",
		"est_duration",
		"cost",
		"additional_cost",
		"additional_cost_desc",
		"unit",
		"price_of_step",
		"status",
	}
)

type SvcPackageDTO struct {
	Id           uuid.UUID  `db:"id"`
	ServiceId    uuid.UUID  `db:"service_id"`
	Name         string     `db:"name"`
	Description  string     `db:"description"`
	ComboDays    int        `db:"combo_days"`
	Discount     int        `db:"discount"`
	TimeInterval int        `db:"time_interval"`
	Status       string     `db:"status"`
	CreatedAt    *time.Time `db:"created_at"`
}

func (dto *SvcPackageDTO) ToSvcPackageEntity() (*svcpackagedomain.ServicePackage, error) {
	return svcpackagedomain.NewServicePackage(
		dto.Id,
		dto.ServiceId,
		dto.Name,
		dto.Description,
		dto.ComboDays,
		dto.Discount,
		dto.TimeInterval,
		svcpackagedomain.EnumSvcPackageStatus(dto.Status),
		dto.CreatedAt,
	)
}

func ToSvcPackageDTO(data *svcpackagedomain.ServicePackage) *SvcPackageDTO {
	return &SvcPackageDTO{
		Id:           data.GetID(),
		ServiceId:    data.GetServiceID(),
		Name:         data.GetName(),
		Description:  data.GetDescription(),
		ComboDays:    data.GetComboDays(),
		Discount:     data.GetDiscount(),
		TimeInterval: data.GetTimeInterVal(),
		Status:       data.GetStatus().String(),
	}
}

type SvcTaskDTO struct {
	Id                 uuid.UUID `db:"id"`
	SvcPackageId       uuid.UUID `db:"service_package_id"`
	IsMustHave         bool      `db:"is_must_have"`
	TaskOrder          int       `db:"task_order"`
	Name               string    `db:"name"`
	Description        string    `db:"description"`
	StaffAdvice        string    `db:"staff_advice"`
	EstDuration        int       `db:"est_duration"`
	Cost               float64   `db:"cost"`
	AdditionalCost     float64   `db:"additional_cost"`
	AdditionalCostDesc string    `db:"additional_cost_desc"`
	Unit               string    `db:"unit"`
	PriceOfStep        int       `db:"price_of_step"`
	Status             string    `db:"status"`
}

func (dto *SvcTaskDTO) ToSvcTaskEntity() (*svcpackagedomain.ServiceTask, error) {
	return svcpackagedomain.NewServiceTask(
		dto.Id,
		dto.SvcPackageId,
		dto.IsMustHave,
		dto.TaskOrder,
		dto.Name,
		dto.Description,
		dto.StaffAdvice,
		dto.EstDuration,
		dto.Cost,
		dto.AdditionalCost,
		dto.AdditionalCostDesc,
		svcpackagedomain.EnumSvcTaskUnit(dto.Unit),
		dto.PriceOfStep,
		svcpackagedomain.EnumSvcTaskStatus(dto.Status),
	)
}

func ToSvcTaskDTO(data *svcpackagedomain.ServiceTask) *SvcTaskDTO {
	return &SvcTaskDTO{
		Id:                 data.GetID(),
		SvcPackageId:       data.GetSvcPackageID(),
		IsMustHave:         data.GetIsMustHave(),
		TaskOrder:          data.GetTaskOrder(),
		Name:               data.GetName(),
		Description:        data.GetDescription(),
		StaffAdvice:        data.GetStaffAdvice(),
		EstDuration:        data.GetEstDuration(),
		Cost:               data.GetCost(),
		AdditionalCost:     data.GetAdditionCost(),
		AdditionalCostDesc: data.GetAdditionCostDesc(),
		Unit:               data.GetUnit().String(),
		PriceOfStep:        data.GetPriceOfStep(),
		Status:             data.GetStatus().String(),
	}
}

type SvcPackageUsageCountDTO struct {
	Id         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	UsageCount int       `db:"usage_count"`
}

func (dto *SvcPackageUsageCountDTO) ToSvcPackageUsageCountEntity() (*svcpackagedomain.ServicePackageUsage, error) {
	return svcpackagedomain.NewServicePackageUsage(
		dto.Id,
		dto.Name,
		dto.UsageCount,
	)
}
