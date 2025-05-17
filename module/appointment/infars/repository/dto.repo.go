package appointmentrepository

import (
	"time"

	"github.com/google/uuid"

	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

var (
	TABLE_APPOINTMENT = `appointments`

	CREATE_APPOINTMENT = []string{
		"id",
		"service_id",
		"customized_package_id",
		"nursing_id",
		"patient_id",
		"patient_address",
		"patient_lat_lng",
		"est_date",
		"act_date",
		"status",
		"total_est_duration",
	}

	GET_APPOINTMENT = []string{
		"id",
		"service_id",
		"customized_package_id",
		"nursing_id",
		"patient_id",
		"patient_address",
		"patient_lat_lng",
		"est_date",
		"act_date",
		"status",
		"total_est_duration",
		"created_at",
	}

	UPDATE_APPOINTMENT = []string{
		"patient_address",
		"patient_lat_lng",
		"nursing_id",
		"act_date",
		"status",
		"total_est_duration",
	}
)

type AppointmentDTO struct {
	Id                  uuid.UUID  `db:"id"`
	ServiceId           uuid.UUID  `db:"service_id"`
	CustomizedPackageId uuid.UUID  `db:"customized_package_id"`
	NursingId           *uuid.UUID `db:"nursing_id"`
	PatientId           uuid.UUID  `db:"patient_id"`
	PatientAddress      string     `db:"patient_address"`
	PatientLagLng       string     `db:"patient_lat_lng"`
	EstDate             time.Time  `db:"est_date"`
	ActDate             *time.Time `db:"act_date"`
	Status              string     `db:"status"`
	TotalEstDuration    int        `db:"total_est_duration"`
	CreatedAt           *time.Time `db:"created_at"`
}

func (dto *AppointmentDTO) ToAppointmentEntity() (*appointmentdomain.Appointment, error) {
	return appointmentdomain.NewAppointment(
		dto.Id,
		dto.ServiceId,
		dto.CustomizedPackageId,
		dto.PatientId,
		dto.NursingId,
		dto.PatientAddress,
		dto.PatientLagLng,
		appointmentdomain.EnumAppointmentStatus(dto.Status),
		dto.TotalEstDuration,
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
		PatientAddress:      data.GetPatientAddress(),
		PatientLagLng:       data.GetPatientLatLng(),
		EstDate:             data.GetEstDate(),
		ActDate:             data.GetActDate(),
		Status:              data.GetStatus().String(),
		TotalEstDuration:    data.GetTotalEstDuration(),
	}
}
