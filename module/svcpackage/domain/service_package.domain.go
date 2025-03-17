package svcpackagedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type ServicePackage struct {
	id           uuid.UUID
	serviceId    uuid.UUID
	name         string
	description  string
	comboDays    int
	discount     int
	timeInterval int
	status       SvcPackageStatus
	createdAt    *time.Time
}

func (a *ServicePackage) GetID() uuid.UUID {
	return a.id
}

func (a *ServicePackage) GetServiceID() uuid.UUID {
	return a.serviceId
}

func (a *ServicePackage) GetName() string {
	return a.name
}

func (a *ServicePackage) GetDescription() string {
	return a.description
}

func (a *ServicePackage) GetComboDays() int {
	return a.comboDays
}

func (a *ServicePackage) GetDiscount() int {
	return a.discount
}

func (a *ServicePackage) GetTimeInterVal() int {
	return a.timeInterval
}

func (a *ServicePackage) GetStatus() SvcPackageStatus {
	return a.status
}

func (a *ServicePackage) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewServicePackage(
	id, serviceId uuid.UUID,
	name, description string,
	comboDays, discount, timeInterval int,
	status SvcPackageStatus,
	createdAt *time.Time,
) (*ServicePackage, error) {
	return &ServicePackage{
		id:           id,
		serviceId:    serviceId,
		name:         name,
		description:  description,
		comboDays:    comboDays,
		discount:     discount,
		timeInterval: timeInterval,
		status:       status,
		createdAt:    createdAt,
	}, nil
}

type SvcPackageStatus int

const (
	SvcPackageStatusAvailable SvcPackageStatus = iota
	SvcPackageStatusUnavailable
)

func (r SvcPackageStatus) String() string {
	switch r {
	case SvcPackageStatusAvailable:
		return "available"
	case SvcPackageStatusUnavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}

func EnumSvcPackageStatus(s string) SvcPackageStatus {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "available":
		return SvcPackageStatusAvailable
	default:
		return SvcPackageStatusUnavailable
	}
}
