package svcpackagequeries

import (
	"time"

	"github.com/google/uuid"

	svcpackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/svcpackage/domain"
)

type ServicePackageDTO struct {
	Id           uuid.UUID `json:"id"`
	ServiceId    uuid.UUID `json:"service-id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	ComboDays    int       `json:"combo-days"`
	Discount     int       `json:"discount"`
	TimeInterval int       `json:"time-interval"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created-at"`
}

func toServicePackageDTO(entity *svcpackagedomain.ServicePackage) *ServicePackageDTO {
	return &ServicePackageDTO{
		Id:           entity.GetID(),
		ServiceId:    entity.GetServiceID(),
		Name:         entity.GetName(),
		Description:  entity.GetDescription(),
		ComboDays:    entity.GetComboDays(),
		Discount:     entity.GetDiscount(),
		TimeInterval: entity.GetTimeInterVal(),
		Status:       entity.GetStatus().String(),
		CreatedAt:    entity.GetCreatedAt(),
	}
}

type ServiceTaskDTO struct {
	Id                 uuid.UUID `json:"id"`
	SvcPackageId       uuid.UUID `json:"svcpackage-id"`
	IsMustHave         bool      `json:"is-must-have"`
	TaskOrder          int       `json:"task-order"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	StaffAdvice        string    `json:"staff-advice"`
	EstDuration        int       `json:"est-duration"`
	Cost               float64   `json:"cost"`
	AdditionalCost     float64   `json:"additional-cost"`
	AdditionalCostDesc string    `json:"additional-cost-desc"`
	Unit               string    `json:"unit"`
	PriceOfStep        int       `json:"price-of-step"`
	Status             string    `json:"status"`
}

func toServiceTaskDTO(entity *svcpackagedomain.ServiceTask) *ServiceTaskDTO {
	return &ServiceTaskDTO{
		Id:                 entity.GetID(),
		SvcPackageId:       entity.GetSvcPackageID(),
		IsMustHave:         entity.GetIsMustHave(),
		TaskOrder:          entity.GetTaskOrder(),
		Name:               entity.GetName(),
		Description:        entity.GetDescription(),
		StaffAdvice:        entity.GetStaffAdvice(),
		EstDuration:        entity.GetEstDuration(),
		Cost:               entity.GetCost(),
		AdditionalCost:     entity.GetAdditionCost(),
		AdditionalCostDesc: entity.GetAdditionCostDesc(),
		Unit:               entity.GetUnit().String(),
		PriceOfStep:        entity.GetPriceOfStep(),
		Status:             entity.GetStatus().String(),
	}
}
