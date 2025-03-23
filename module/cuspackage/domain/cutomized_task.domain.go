package cuspackagedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type CustomizedTask struct {
	id           uuid.UUID
	svcTaskid    uuid.UUID
	cusPackageId uuid.UUID
	taskOrder    int
	name         string
	clientNote   string
	staffAdvice  string
	estDuration  int
	totalCost    float64
	unit         CusTaskUnit
	totalUnit    int
	estDate      time.Time
	actDate      time.Time
	status       CusTaskStatus
}

func (a *CustomizedTask) GetID() uuid.UUID {
	return a.id
}

func (a *CustomizedTask) GetSvcTaskID() uuid.UUID {
	return a.svcTaskid
}

func (a *CustomizedTask) GetServicePlanCustomID() uuid.UUID {
	return a.cusPackageId
}

func (a *CustomizedTask) GetTaskOrder() int {
	return a.taskOrder
}

func (a *CustomizedTask) GetName() string {
	return a.name
}

func (a *CustomizedTask) GetClientNote() string {
	return a.clientNote
}

func (a *CustomizedTask) GetStaffAdvice() string {
	return a.staffAdvice
}

func (a *CustomizedTask) GetEstDuration() int {
	return a.estDuration
}

func (a *CustomizedTask) GetTotalCost() float64 {
	return a.totalCost
}

func (a *CustomizedTask) GetUnit() CusTaskUnit {
	return a.unit
}

func (a *CustomizedTask) GetTotalUnit() int {
	return a.totalUnit
}

func (a *CustomizedTask) GetEstDate() time.Time {
	return a.estDate
}

func (a *CustomizedTask) GetActDate() time.Time {
	return a.actDate
}

func (a *CustomizedTask) GetStatus() CusTaskStatus {
	return a.status
}

func NewCustomizedTask(
	id, svcTaskId, cusPackageId uuid.UUID,
	isMustHave bool,
	taskOrder int,
	name, clientNote, staffAdvice string,
	estDuration int,
	totalCost float64,
	unit CusTaskUnit,
	totalUnit int,
	estDate, actDate time.Time,
	status CusTaskStatus,
) (*CustomizedTask, error) {
	return &CustomizedTask{
		id:           id,
		svcTaskid:    svcTaskId,
		cusPackageId: cusPackageId,
		taskOrder:    taskOrder,
		name:         name,
		clientNote:   clientNote,
		staffAdvice:  staffAdvice,
		estDuration:  estDuration,
		totalCost:    totalCost,
		unit:         unit,
		totalUnit:    totalUnit,
		estDate:      estDate,
		actDate:      actDate,
		status:       status,
	}, nil
}

type (
	CusTaskStatus int
	CusTaskUnit   int
)

const (
	CusTaskStatusNotDone CusTaskStatus = iota
	CusTaskStatusDone
)

func (r CusTaskStatus) String() string {
	switch r {
	case CusTaskStatusDone:
		return "done"
	case CusTaskStatusNotDone:
		return "not_done"
	default:
		return "unknown"
	}
}

func EnumCusTaskStatus(s string) CusTaskStatus {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "done":
		return CusTaskStatusDone
	default:
		return CusTaskStatusNotDone
	}
}

const (
	CusTaskUnitQuantity CusTaskUnit = iota
	CusTaskUnitTime
)

func (r CusTaskUnit) String() string {
	switch r {
	case CusTaskUnitQuantity:
		return "quantity"
	case CusTaskUnitTime:
		return "time"
	default:
		return "unknown"
	}
}

func EnumCusTaskUnit(s string) CusTaskUnit {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "quantity":
		return CusTaskUnitQuantity
	default:
		return CusTaskUnitTime
	}
}
