package svcplancustomdomain

import (
	"time"

	"github.com/google/uuid"
)

type ServicePlanCustom struct {
	id            uuid.UUID
	serviceId     uuid.UUID
	appointmentId uuid.UUID
	name          string
	beginDate     time.Time
	status        SvcPlanCusStatus
	createdAt     *time.Time
}

func (a *ServicePlanCustom) GetID() uuid.UUID {
	return a.id
}

func (a *ServicePlanCustom) GetServiceID() uuid.UUID {
	return a.serviceId
}

func (a *ServicePlanCustom) GetAppointmentID() uuid.UUID {
	return a.appointmentId
}

func (a *ServicePlanCustom) GetName() string {
	return a.name
}

func (a *ServicePlanCustom) GetBeginDate() time.Time {
	return a.beginDate
}

func (a *ServicePlanCustom) GetStatus() SvcPlanCusStatus {
	return a.status
}

func (a *ServicePlanCustom) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewServicePlanCustom(id, serviceId, appointmentId uuid.UUID, name string, beginDate time.Time, status SvcPlanCusStatus, createdAt *time.Time) (*ServicePlanCustom, error) {
	return &ServicePlanCustom{
		id:        id,
		serviceId: serviceId,
		name:      name,
		beginDate: beginDate,
		status:    status,
		createdAt: createdAt,
	}, nil
}

type SvcPlanCusStatus int

const (
	SvcPlanStatusAvailable SvcPlanCusStatus = iota
	SvcPlanStatusUnavailable
)

func (r SvcPlanCusStatus) String() string {
	switch r {
	case SvcPlanStatusAvailable:
		return "available"
	case SvcPlanStatusUnavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}
