package svcpackagedomain

import (
	"strings"

	"github.com/google/uuid"
)

type ServiceTask struct {
	id                 uuid.UUID
	svcPackageId       uuid.UUID
	isMustHave         bool
	taskOrder          int
	name               string
	description        string
	staffAdvice        string
	estDuration        int
	cost               float64
	additionalCost     float64
	additionalCostDesc string
	unit               SvcTaskUnit
	priceOfStep        int
	status             SvcTaskStatus
}

func (a *ServiceTask) GetID() uuid.UUID {
	return a.id
}

func (a *ServiceTask) GetSvcPackageID() uuid.UUID {
	return a.svcPackageId
}

func (a *ServiceTask) GetIsMustHave() bool {
	return a.isMustHave
}

func (a *ServiceTask) GetTaskOrder() int {
	return a.taskOrder
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
	return a.additionalCost
}

func (a *ServiceTask) GetAdditionCostDesc() string {
	return a.additionalCostDesc
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
	id, svcPackageId uuid.UUID,
	isMustHave bool,
	taskOrder int,
	name, description, staffAdvice string,
	estDuration int,
	cost, additionCost float64,
	additionCostDesc string,
	unit SvcTaskUnit,
	priceOfStep int,
	status SvcTaskStatus,
) (*ServiceTask, error) {
	return &ServiceTask{
		id:                 id,
		svcPackageId:       svcPackageId,
		isMustHave:         isMustHave,
		taskOrder:          taskOrder,
		name:               name,
		description:        description,
		staffAdvice:        staffAdvice,
		estDuration:        estDuration,
		cost:               cost,
		additionalCost:     additionCost,
		additionalCostDesc: additionCostDesc,
		unit:               unit,
		priceOfStep:        priceOfStep,
		status:             status,
	}, nil
}

type (
	SvcTaskStatus int
	SvcTaskUnit   int
)

const (
	SvcTaskStatusAvailable SvcTaskStatus = iota
	SvcTaskStatusUnavailable
)

func (r SvcTaskStatus) String() string {
	switch r {
	case SvcTaskStatusAvailable:
		return "available"
	case SvcTaskStatusUnavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}

func EnumSvcTaskStatus(s string) SvcTaskStatus {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "available":
		return SvcTaskStatusAvailable
	default:
		return SvcTaskStatusUnavailable
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

func EnumSvcTaskUnit(s string) SvcTaskUnit {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "quantity":
		return SvcTaskUnitQuantity
	default:
		return SvcTaskUnitTime
	}
}
