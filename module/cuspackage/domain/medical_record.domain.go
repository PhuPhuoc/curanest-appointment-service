package cuspackagedomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	id            uuid.UUID
	cusTaskId     uuid.UUID
	nursingId     *uuid.UUID
	nursingReport string
	staffAdvice   string
	status        string
	createdAt     *time.Time
}

func NewMedicalRecord(
	id, cusTaskId uuid.UUID,
	nursingId *uuid.UUID,
	nursingReport, staffAdvice string,
	status string,
	createdAt *time.Time,
) (*MedicalRecord, error) {
	return &MedicalRecord{
		id:            id,
		cusTaskId:     cusTaskId,
		nursingId:     nursingId,
		nursingReport: nursingReport,
		staffAdvice:   staffAdvice,
		status:        status,
		createdAt:     createdAt,
	}, nil
}

func (a *MedicalRecord) GetID() uuid.UUID {
	return a.id
}

func (a *MedicalRecord) GetCusTaskId() uuid.UUID {
	return a.cusTaskId
}

func (a *MedicalRecord) GetNursingId() *uuid.UUID {
	return a.nursingId
}

func (a *MedicalRecord) GetNursingReport() string {
	return a.nursingReport
}

func (a *MedicalRecord) GetStaffAdvice() string {
	return a.staffAdvice
}

func (a *MedicalRecord) GetStatus() string {
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
