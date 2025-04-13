package appointmentdomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	id               uuid.UUID
	serviceId        uuid.UUID
	cusPackageId     uuid.UUID
	nursingId        *uuid.UUID
	patientId        uuid.UUID
	estDate          time.Time
	actDate          *time.Time
	status           AppointmentStatus
	totalEstDuration int
	createdAt        *time.Time
}

func (a *Appointment) GetID() uuid.UUID {
	return a.id
}

func (a *Appointment) GetServiceID() uuid.UUID {
	return a.serviceId
}

func (a *Appointment) GetCusPackageID() uuid.UUID {
	return a.cusPackageId
}

func (a *Appointment) GetNursingID() *uuid.UUID {
	return a.nursingId
}

func (a *Appointment) GetPatientID() uuid.UUID {
	return a.patientId
}

func (a *Appointment) GetStatus() AppointmentStatus {
	return a.status
}

func (a *Appointment) GetTotalEstDuration() int {
	return a.totalEstDuration
}

func (a *Appointment) GetEstDate() time.Time {
	return a.estDate
}

func (a *Appointment) GetActDate() *time.Time {
	return a.actDate
}

func (a *Appointment) GetCreatedAt() *time.Time {
	return a.createdAt
}

func NewAppointment(
	id, serviceId, cusPackageId, patientId uuid.UUID,
	nursingId *uuid.UUID,
	status AppointmentStatus,
	totalEstDuration int,
	estDate time.Time,
	actDate *time.Time,
	createdAt *time.Time,
) (*Appointment, error) {
	return &Appointment{
		id:               id,
		serviceId:        serviceId,
		cusPackageId:     cusPackageId,
		nursingId:        nursingId,
		patientId:        patientId,
		status:           status,
		totalEstDuration: totalEstDuration,
		estDate:          estDate,
		actDate:          actDate,
		createdAt:        createdAt,
	}, nil
}

type AppointmentStatus int

const (
	AppStatusWaiting AppointmentStatus = iota
	AppStatusSuccess
	AppStatusConfirmed
	AppStatusRefused
	AppStatusChanged
)

func EnumAppointmentStatus(s string) AppointmentStatus {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "success":
		return AppStatusSuccess
	case "confirmed":
		return AppStatusConfirmed
	case "refused":
		return AppStatusRefused
	case "changed":
		return AppStatusChanged
	default:
		return AppStatusWaiting
	}
}

func (r AppointmentStatus) String() string {
	switch r {
	case AppStatusWaiting:
		return "waiting"
	case AppStatusSuccess:
		return "success"
	case AppStatusConfirmed:
		return "confirmed"
	case AppStatusRefused:
		return "refused"
	case AppStatusChanged:
		return "changed"
	default:
		return "unknown"
	}
}
