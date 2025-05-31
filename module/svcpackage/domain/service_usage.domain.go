package svcpackagedomain

import (
	"github.com/google/uuid"
)

type ServicePackageUsage struct {
	id         uuid.UUID
	name       string
	usageCount int
}

func (a *ServicePackageUsage) GetID() uuid.UUID {
	return a.id
}

func (a *ServicePackageUsage) GetName() string {
	return a.name
}

func (a *ServicePackageUsage) GetUsageCount() int {
	return a.usageCount
}

func NewServicePackageUsage(
	id uuid.UUID,
	name string,
	usageCount int,
) (*ServicePackageUsage, error) {
	return &ServicePackageUsage{
		id:         id,
		name:       name,
		usageCount: usageCount,
	}, nil
}
