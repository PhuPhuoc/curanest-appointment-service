package cuspackagecommands

import (
	"time"

	"github.com/google/uuid"
)

type ReqCreatePackageTaskDTO struct {
	Dates          []time.Time               `json:"dates"`
	PatientId      uuid.UUID                 `json:"patient-id"`
	PatientAddress string                    `json:"patient-address"`
	PatientLatLng  string                    `json:"-"`
	NursingId      *uuid.UUID                `json:"nursing-id"`
	SvcPackageId   uuid.UUID                 `json:"svcpackage-id" binding:"required"`
	TaskInfos      []CreateCustomizedTaskDTO `json:"task-infos"`
}

type CreateCustomizedTaskDTO struct {
	SvcTaskId   uuid.UUID `json:"svctask-id" binding:"required"`
	ClientNote  string    `json:"client-note"`
	TotalCost   float64   `json:"total-cost"`
	TotalUnit   int       `json:"total-unit"`
	EstDuration int       `json:"est-duration"`
}

// Location chứa lat và lng
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Geometry chứa location
type Geometry struct {
	Location Location `json:"location"`
}

// Result đại diện cho mỗi phần tử trong mảng results
type GoongResult struct {
	Geometry Geometry `json:"geometry"`
}

// APIResponse đại diện cho toàn bộ response từ API
type GoongAPIResponse struct {
	Results []GoongResult `json:"results"`
	Status  string        `json:"status"`
}
