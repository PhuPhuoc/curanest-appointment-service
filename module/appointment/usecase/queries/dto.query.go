package appointmentqueries

import (
	"time"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type FilterDashboardDTO struct {
	IsAdmin    bool       `json:"is-admin"`
	CategoryId *uuid.UUID `json:"category-id,omitempty"`
	DateFrom   string     `json:"date-from,omitempty"`
	DateTo     string     `json:"date-to,omitempty"`
}

type ItemDashboardDTO struct {
	TotalService int `json:"total-service"`
	UpcomingApps int `json:"upcoming-apps"`
	WaitingApps  int `json:"waiting-apps"`
	TotalApps    int `json:"total-apps"`
}

type FilterGetNursingTimesheetDTO struct {
	NursingId   uuid.UUID  `json:"nursing-id,omitempty"`
	EstDateFrom *time.Time `json:"est-date-from,omitempty"`
	EstDateTo   *time.Time `json:"est-date-to,omitempty"`
}

type TimesheetDTO struct {
	AppointmentId    uuid.UUID `json:"appointment-id"`
	PatientId        uuid.UUID `json:"patient-id"`
	EstDate          time.Time `json:"est-date"`
	EstEndTime       time.Time `json:"est-end-time"`
	Status           string    `json:"status"`
	TotalEstDuration int       `json:"total-est-duration"`
	EstTravelTime    int       `json:"est-travel-time"`
}

func toTimesheetDTO(data *appointmentdomain.Appointment, estTravelTime int) *TimesheetDTO {
	estDate := data.GetEstDate()
	endTime := estDate.Add(time.Duration(data.GetTotalEstDuration()+estTravelTime) * time.Minute)
	dto := &TimesheetDTO{
		AppointmentId:    data.GetID(),
		PatientId:        data.GetPatientID(),
		EstDate:          data.GetEstDate(),
		EstEndTime:       endTime,
		Status:           data.GetStatus().String(),
		TotalEstDuration: data.GetTotalEstDuration(),
		EstTravelTime:    estTravelTime,
	}
	return dto
}

type FilterGetAppointmentDTO struct {
	Id                *uuid.UUID                           `json:"id,omitempty"`
	CategoryId        *uuid.UUID                           `json:"category-id,omitempty"`
	ServiceId         *uuid.UUID                           `json:"service-id,omitempty"`
	CusPackageId      *uuid.UUID                           `json:"cuspackage-id,omitempty"`
	NursingId         *uuid.UUID                           `json:"nursing-id,omitempty"`
	PatientId         *uuid.UUID                           `json:"patient-id,omitempty"`
	HadNurse          *bool                                `json:"had-nurse,omitempty"`
	AppointmentStatus *appointmentdomain.AppointmentStatus `json:"appointment-status,omitempty" binding:"oneof=success waiting confirmed refused change"`
	EstDateFrom       *time.Time                           `json:"est-date-from,omitempty"`
	EstDateTo         *time.Time                           `json:"est-date-to,omitempty"`
	ApplyPaging       *bool                                `json:"apply-paging,omitempty"`
	Paging            *common.Paging                       `json:"-"`
}

type AppointmentDTO struct {
	Id               uuid.UUID  `json:"id"`
	ServiceId        uuid.UUID  `json:"service-id"`
	CusPackageId     uuid.UUID  `json:"cuspackage-id"`
	NursingId        *uuid.UUID `json:"nursing-id"`
	PatientId        uuid.UUID  `json:"patient-id"`
	PatientAddress   string     `json:"patient-address"`
	PatientLatLng    string     `json:"patient-lat-lng"`
	EstDate          time.Time  `json:"est-date"`
	ActDate          *time.Time `json:"act-date"`
	Status           string     `json:"status"`
	IsPaid           bool       `json:"is-paid"`
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
		PatientAddress:   data.GetPatientAddress(),
		PatientLatLng:    data.GetPatientLatLng(),
		EstDate:          data.GetEstDate(),
		ActDate:          data.GetActDate(),
		Status:           data.GetStatus().String(),
		TotalEstDuration: data.GetTotalEstDuration(),
		CreatedAt:        data.GetCreatedAt(),
	}
	return dto
}

func (dto AppointmentDTO) ToAppointmentDomain() (*appointmentdomain.Appointment, error) {
	return appointmentdomain.NewAppointment(
		dto.Id,
		dto.ServiceId,
		dto.CusPackageId,
		dto.PatientId,
		dto.NursingId,
		dto.PatientAddress,
		dto.PatientLatLng,
		appointmentdomain.EnumAppointmentStatus(dto.Status),
		dto.TotalEstDuration,
		dto.EstDate,
		dto.ActDate,
		dto.CreatedAt,
	)
}

type NurseDateMapping struct {
	NurseId      uuid.UUID `json:"nurse-id"`
	EstStartDate time.Time `json:"est-start-date"`
	EstDuration  int       `json:"est-duration"`
}

type CheckNursesAvailabilityRequestDTO struct {
	NursesDates []NurseDateMapping `json:"nurses-dates"`
}

type NurseDateMappingResult struct {
	NurseId        uuid.UUID `json:"nurse-id"`
	EstStartDate   time.Time `json:"est-start-date"`
	EstDuration    int       `json:"est-duration"`
	IsAvailability bool      `json:"is-availability"`
}

type NurseDTO struct {
	NurseId          uuid.UUID `json:"nurse-id"`
	NursePicture     string    `json:"nurse-picture"`
	NurseName        string    `json:"nurse-name"`
	Gender           bool      `json:"gender"`
	CurrentWorkPlace string    `json:"current-work-place"`
	Rate             float64   `json:"rate"`
}
