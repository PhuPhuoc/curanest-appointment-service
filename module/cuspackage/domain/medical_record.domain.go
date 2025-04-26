package cuspackagedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	id            uuid.UUID
	nursingId     *uuid.UUID
	appointmentId uuid.UUID
	nursingReport string
	staffConfirm  string
	status        RecordStatus
	createdAt     *time.Time
}

func NewMedicalRecord(
	id, appointmentId uuid.UUID,
	nursingId *uuid.UUID,
	nursingReport, staffConfirm string,
	status RecordStatus,
	createdAt *time.Time,
) (*MedicalRecord, error) {
	return &MedicalRecord{
		id:            id,
		appointmentId: appointmentId,
		nursingId:     nursingId,
		nursingReport: nursingReport,
		staffConfirm:  staffConfirm,
		status:        status,
		createdAt:     createdAt,
	}, nil
}

func (a *MedicalRecord) GetID() uuid.UUID {
	return a.id
}

func (a *MedicalRecord) GetAppointmentId() uuid.UUID {
	return a.appointmentId
}

func (a *MedicalRecord) GetNursingId() *uuid.UUID {
	return a.nursingId
}

func (a *MedicalRecord) GetNursingReport() string {
	return a.nursingReport
}

func (a *MedicalRecord) GetStaffConfirm() string {
	return a.staffConfirm
}

func (a *MedicalRecord) GetStatus() RecordStatus {
	return a.status
}

func (a *MedicalRecord) GetCreatedAt() *time.Time {
	return a.createdAt
}

type RecordStatus int

const (
	RecordStatusNotDone RecordStatus = iota
	RecordStatusDone
)

func EnumRecordStatus(s string) RecordStatus {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "done":
		return RecordStatusDone
	default:
		return RecordStatusNotDone
	}
}

func (r RecordStatus) String() string {
	switch r {
	case RecordStatusNotDone:
		return "not_done"
	case RecordStatusDone:
		return "done"
	default:
		return "unknown"
	}
}
