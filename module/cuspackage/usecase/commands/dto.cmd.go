package cuspackagecommands

import (
	"time"

	"github.com/google/uuid"
)

type AddMoreCustaskRequestDTO struct {
	AppointmentId uuid.UUID                 `json:"appointment-id" binding:"required"`
	CuspackageId  uuid.UUID                 `json:"cuspackage-id" binding:"required"`
	TaskInfos     []CreateCustomizedTaskDTO `json:"task-infos"`
}

type DateNursingMapping struct {
	Date      time.Time  `json:"date"`
	NursingId *uuid.UUID `json:"nursing-id"`
}

type ReqCreatePackageTaskDTO struct {
	DateNurseMappings []DateNursingMapping      `json:"date-nurse-mappings"`
	PatientId         uuid.UUID                 `json:"patient-id"`
	PatientAddress    string                    `json:"patient-address"`
	PatientLatLng     string                    `json:"-"`
	SvcPackageId      uuid.UUID                 `json:"svcpackage-id" binding:"required"`
	TaskInfos         []CreateCustomizedTaskDTO `json:"task-infos"`
}

type CreateCustomizedTaskDTO struct {
	SvcTaskId   uuid.UUID `json:"svctask-id" binding:"required"`
	ClientNote  string    `json:"client-note"`
	TotalCost   float64   `json:"total-cost"`
	TotalUnit   int       `json:"total-unit"`
	EstDuration int       `json:"est-duration"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type GoongResult struct {
	Geometry Geometry `json:"geometry"`
}

type GoongAPIResponse struct {
	Results []GoongResult `json:"results"`
	Status  string        `json:"status"`
}

type UpdateMedicalRecordDTO struct {
	NursingReport     *string `json:"nursing-report"`
	StaffConfirmation *string `json:"staff-confirmation"`
}
