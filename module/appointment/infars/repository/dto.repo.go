package appointmentrepository

import (
	"time"

	"github.com/google/uuid"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

var (
	TABLE_CUSPACKAGE = `customized_packages`

	CREATE_CUSPACKAGE = []string{
		"id",
		"service_id",
		"customized_package_id",
		"nursing_id",
		"patient_id",
		"est_date",
		"act_date",
		"status",
	}

	GET_CUSPACKAGE = []string{
		"id",
		"service_id",
		"customized_package_id",
		"nursing_id",
		"patient_id",
		"est_date",
		"act_date",
		"created_at",
	}

	UPDATE_CUSPACKAGE = []string{
		"nursing_id",
		"act_date",
	}
)

type AppointmentDTO struct {
	Id                  uuid.UUID  `db:"id"`
	ServiceId           uuid.UUID  `db:"service_package_id"`
	CustomizedPackageId uuid.UUID  `db:"patient_id"`
	NursingId           *uuid.UUID `db:"patient_id"`
	PatientId           uuid.UUID  `db:"patient_id"`
	EstDate             time.Time  `db:"created_at"`
	ActDate             *time.Time `db:"created_at"`
	Status              string     `db:"created_at"`
	CreatedAt           *time.Time `db:"created_at"`
}

func (dto *AppointmentDTO) ToAppointmentEntity() (*appointmentdomain.Appointment, error) {
	return appointmentdomain.NewAppointment(
		dto.Id,
		dto.ServiceId,
		dto.CustomizedPackageId,
		dto.PatientId,
		dto.NursingId,
		appointmentdomain.EnumAppointmentStatus(dto.Status),
		dto.EstDate,
		dto.ActDate,
		dto.CreatedAt,
	)
}

func ToAppointmentDTO(data *appointmentdomain.Appointment) *AppointmentDTO {
	return &AppointmentDTO{
		Id:                  data.GetID(),
		ServiceId:           data.GetServiceID(),
		CustomizedPackageId: data.GetCusPackageID(),
		PatientId:           data.GetPatientID(),
		NursingId:           data.GetNursingID(),
		EstDate:             data.GetEstDate(),
		ActDate:             data.GetActDate(),
		Status:              data.Status.String(),
	}
}
