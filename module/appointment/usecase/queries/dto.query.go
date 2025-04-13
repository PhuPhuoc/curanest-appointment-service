package appointmentqueries

import (
	"time"

	"github.com/google/uuid"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type FilterGetAppointmentDTO struct {
	ServiceId         *uuid.UUID                           `json:"service-id,omitempty"`
	CusPackageId      *uuid.UUID                           `json:"cuspackage-id,omitempty"`
	NursingId         *uuid.UUID                           `json:"nursing-id,omitempty"`
	PatientId         *uuid.UUID                           `json:"patient-id,omitempty"`
	HadNurse          *bool                                `json:"had-nurse,omitempty"`
	AppointmentStatus *appointmentdomain.AppointmentStatus `json:"appointment-status,omitempty" binding:"oneof=success waiting confirmed refused change"`
	EstDateFrom       *time.Time                           `json:"est-date-from,omitempty"`
	EstDateTo         *time.Time                           `json:"est-date-to,omitempty"`
}

type AppointmentDTO struct {
	Id               uuid.UUID  `json:"id"`
	ServiceId        uuid.UUID  `json:"service-id"`
	CusPackageId     uuid.UUID  `json:"cuspackage-id"`
	NursingId        *uuid.UUID `json:"nursing-id"`
	PatientId        uuid.UUID  `json:"patient-id"`
	EstDate          time.Time  `json:"est-date"`
	ActDate          *time.Time `json:"act-date"`
	Status           string     `json:"status"`
	TotalEstDuration int        `json:"total-est-duration"`
	CreatedAt        *time.Time `json:"created-at"`
}

func toAppointmentDTO(data *appointmentdomain.Appointment) *AppointmentDTO {
	dto := &AppointmentDTO{
		Id:               data.GetID(),
		ServiceId:        data.GetServiceID(),
		CusPackageId:     data.GetCusPackageID(),
		NursingId:        data.GetNursingID(),
		PatientId:        data.GetPatientID(),
		EstDate:          data.GetEstDate(),
		ActDate:          data.GetActDate(),
		Status:           data.GetStatus().String(),
		TotalEstDuration: data.GetTotalEstDuration(),
		CreatedAt:        data.GetCreatedAt(),
	}
	return dto
}
