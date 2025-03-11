package svcplancustomdomain

import (
	"time"

	"github.com/google/uuid"
)

type ServiceTaskCustom struct {
	id                  uuid.UUID
	servicePlanCustomId uuid.UUID
	order               int
	name                string
	clientNote          string
	staffAdvice         string
	estDuration         int
	totalCost           float64
	unit                SvcTaskCusUnit
	totalUnit           int
	estDate             time.Time
	actDate             time.Time
	status              SvcTaskCusStatus
}

func (a *ServiceTaskCustom) GetID() uuid.UUID {
	return a.id
}

func (a *ServiceTaskCustom) GetServicePlanCustomID() uuid.UUID {
	return a.servicePlanCustomId
}

func (a *ServiceTaskCustom) GetOrder() int {
	return a.order
}

func (a *ServiceTaskCustom) GetName() string {
	return a.name
}

func (a *ServiceTaskCustom) GetClientNote() string {
	return a.clientNote
}

func (a *ServiceTaskCustom) GetStaffAdvice() string {
	return a.staffAdvice
}

func (a *ServiceTaskCustom) GetEstDuration() int {
	return a.estDuration
}

func (a *ServiceTaskCustom) GetTotalCost() float64 {
	return a.totalCost
}

func (a *ServiceTaskCustom) GetUnit() SvcTaskCusUnit {
	return a.unit
}

func (a *ServiceTaskCustom) GetTotalUnit() int {
	return a.totalUnit
}

func (a *ServiceTaskCustom) GetEstDate() time.Time {
	return a.estDate
}

func (a *ServiceTaskCustom) GetActDate() time.Time {
	return a.actDate
}

func (a *ServiceTaskCustom) GetStatus() SvcTaskCusStatus {
	return a.status
}

func NewServiceTaskCustom(
	id, servicePlanCusId uuid.UUID,
	isMustHave bool,
	order int,
	name, clientNote, staffAdvice string,
	estDuration int,
	totalCost float64,
	unit SvcTaskCusUnit,
	totalUnit int,
	estDate, actDate time.Time,
	status SvcTaskCusStatus,
) (*ServiceTaskCustom, error) {
	return &ServiceTaskCustom{
		id:                  id,
		servicePlanCustomId: servicePlanCusId,
		order:               order,
		name:                name,
		clientNote:          clientNote,
		staffAdvice:         staffAdvice,
		estDuration:         estDuration,
		totalCost:           totalCost,
		unit:                unit,
		totalUnit:           totalUnit,
		estDate:             estDate,
		actDate:             actDate,
		status:              status,
	}, nil
}

type (
	SvcTaskCusStatus int
	SvcTaskCusUnit   int
)

const (
	SvcTaskCusStatusAvailable SvcTaskCusStatus = iota
	SvctaskStatusunavailable
)

func (r SvcTaskCusStatus) String() string {
	switch r {
	case SvcTaskCusStatusAvailable:
		return "available"
	case SvctaskStatusunavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}

const (
	SvcTaskCusUnitQuantity SvcTaskCusUnit = iota
	SvcTaskCusUnitTime
)

func (r SvcTaskCusUnit) String() string {
	switch r {
	case SvcTaskCusUnitQuantity:
		return "quantity"
	case SvcTaskCusUnitTime:
		return "time"
	default:
		return "unknown"
	}
}
