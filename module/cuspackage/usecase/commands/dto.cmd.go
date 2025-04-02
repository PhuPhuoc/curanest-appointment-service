package cuspackagecommands

import (
	"time"

	"github.com/google/uuid"
)

type ReqCreatePackageTaskDTO struct {
	Dates       []time.Time                `json:"dates"`
	PatientId   uuid.UUID                  `json:"patient-id"`
	NursingId   *uuid.UUID                 `json:"nursing-id"`
	PackageInfo CreateCustomizedPackageDTO `json:"package-info"`
	TaskInfos   []CreateCustomizedTaskDTO  `json:"task-infos"`
}

type CreateCustomizedPackageDTO struct {
	SvcPackageId uuid.UUID `json:"svcpackage-id" binding:"required"`
	PatientId    uuid.UUID `json:"patient-id" binding:"required"`
}

type CreateCustomizedTaskDTO struct {
	SvcTaskId   uuid.UUID `json:"svctask-id" binding:"required"`
	ClientNote  string    `json:"client-note"`
	TotalCost   float64   `json:"total-cost"`
	TotalUnit   int       `json:"total-unit"`
	EstDuration int       `json:"est-duration"`
}
