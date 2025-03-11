package svcplandomain

import (
	"time"

	"github.com/google/uuid"
)

type ServicePlan struct {
	id          uuid.UUID
	serviceId   uuid.UUID
	name        string
	description string
	comboDays   int
	discount    int
	status      SvcPlanStatus
	createdAt   *time.Time
}

func (a *ServicePlan) GetID() uuid.UUID {
	return a.id
}

func (a *ServicePlan) GetServiceID() uuid.UUID {
	return a.serviceId
}

func (a *ServicePlan) GetName() string {
	return a.name
}

func (a *ServicePlan) GetDescription() string {
	return a.description
}

func (a *ServicePlan) GetComboDays() int {
	return a.comboDays
}

func (a *ServicePlan) GetDiscount() int {
	return a.discount
}

func (a *ServicePlan) GetStatus() SvcPlanStatus {
	return a.status
}

func (a *ServicePlan) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewServicePlan(
	id, serviceId uuid.UUID,
	name, description string,
	comboDays, discount int,
	status SvcPlanStatus,
	createdAt *time.Time,
) (*ServicePlan, error) {
	return &ServicePlan{
		id:          id,
		serviceId:   serviceId,
		name:        name,
		description: description,
		comboDays:   comboDays,
		discount:    discount,
		status:      status,
		createdAt:   createdAt,
	}, nil
}

type SvcPlanStatus int

const (
	SvcPlanStatusAvailable SvcPlanStatus = iota
	SvcPlanStatusUnavailable
)

func (r SvcPlanStatus) String() string {
	switch r {
	case SvcPlanStatusAvailable:
		return "available"
	case SvcPlanStatusUnavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}
