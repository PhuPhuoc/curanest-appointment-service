package svcplandomain

import (
	"github.com/google/uuid"
)

type ServiceTask struct {
	id               uuid.UUID
	servicePlanId    uuid.UUID
	isMustHave       bool
	order            int
	name             string
	description      string
	staffAdvice      string
	estDuration      int
	cost             float64
	additionCost     float64
	additionCostDesc string
	unit             SvcTaskUnit
	priceOfStep      int
	status           SvcTaskStatus
}

func (a *ServiceTask) GetID() uuid.UUID {
	return a.id
}

func (a *ServiceTask) GetServicePlanID() uuid.UUID {
	return a.servicePlanId
}

func (a *ServiceTask) GetIsMustHave() bool {
	return a.isMustHave
}

func (a *ServiceTask) GetOrder() int {
	return a.order
}

func (a *ServiceTask) GetName() string {
	return a.name
}

func (a *ServiceTask) GetDescription() string {
	return a.description
}

func (a *ServiceTask) GetStaffAdvice() string {
	return a.staffAdvice
}

func (a *ServiceTask) GetEstDuration() int {
	return a.estDuration
}

func (a *ServiceTask) GetCost() float64 {
	return a.cost
}

func (a *ServiceTask) GetAdditionCost() float64 {
	return a.additionCost
}

func (a *ServiceTask) GetAdditionCostDesc() string {
	return a.additionCostDesc
}

func (a *ServiceTask) GetUnit() SvcTaskUnit {
	return a.unit
}

func (a *ServiceTask) GetPriceOfStep() int {
	return a.priceOfStep
}

func (a *ServiceTask) GetStatus() SvcTaskStatus {
	return a.status
}

func NewServiceTask(
	id, servicePlanId uuid.UUID,
	isMustHave bool,
	order int,
	name, description, staffAdvice string,
	estDuration int,
	cost, additionCost float64,
	additionCostDesc string,
	unit SvcTaskUnit,
	priceOfStep int,
	status SvcTaskStatus,
) (*ServiceTask, error) {
	return &ServiceTask{
		id:               id,
		servicePlanId:    servicePlanId,
		isMustHave:       isMustHave,
		order:            order,
		name:             name,
		description:      description,
		staffAdvice:      staffAdvice,
		estDuration:      estDuration,
		cost:             cost,
		additionCost:     additionCost,
		additionCostDesc: additionCostDesc,
		unit:             unit,
		priceOfStep:      priceOfStep,
		status:           status,
	}, nil
}

type (
	SvcTaskStatus int
	SvcTaskUnit   int
)

const (
	SvcTaskStatusAvailable SvcTaskStatus = iota
	SvctaskStatusunavailable
)

func (r SvcTaskStatus) String() string {
	switch r {
	case SvcTaskStatusAvailable:
		return "available"
	case SvctaskStatusunavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}

const (
	SvcTaskUnitQuantity SvcTaskUnit = iota
	SvcTaskUnitTime
)

func (r SvcTaskUnit) String() string {
	switch r {
	case SvcTaskUnitQuantity:
		return "quantity"
	case SvcTaskUnitTime:
		return "time"
	default:
		return "unknown"
	}
}
