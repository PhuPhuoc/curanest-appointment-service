package cuspackagedomain

import (
	"time"

	"github.com/google/uuid"
)

type CustomizedPackage struct {
	id           uuid.UUID
	svcPackageId uuid.UUID
	patientId    *uuid.UUID
	name         string
	beginDate    time.Time
	status       CusPackageStatus
	createdAt    *time.Time
}

func (a *CustomizedPackage) GetID() uuid.UUID {
	return a.id
}

func (a *CustomizedPackage) GetServiceID() uuid.UUID {
	return a.svcPackageId
}

func (a *CustomizedPackage) GetPatientID() *uuid.UUID {
	return a.patientId
}

func (a *CustomizedPackage) GetName() string {
	return a.name
}

func (a *CustomizedPackage) GetBeginDate() time.Time {
	return a.beginDate
}

func (a *CustomizedPackage) GetStatus() CusPackageStatus {
	return a.status
}

func (a *CustomizedPackage) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewCustomizedPackage(id, svcPackageId uuid.UUID, patientId *uuid.UUID, name string, beginDate time.Time, status CusPackageStatus, createdAt *time.Time) (*CustomizedPackage, error) {
	return &CustomizedPackage{
		id:           id,
		svcPackageId: svcPackageId,
		patientId:    patientId,
		name:         name,
		beginDate:    beginDate,
		status:       status,
		createdAt:    createdAt,
	}, nil
}

type CusPackageStatus int

const (
	CusPackageStatusAvailable CusPackageStatus = iota
	CusPackageStatusUnavailable
)

func (r CusPackageStatus) String() string {
	switch r {
	case CusPackageStatusAvailable:
		return "available"
	case CusPackageStatusUnavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}
