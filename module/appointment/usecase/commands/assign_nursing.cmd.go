package apppointmentcommands

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
)

type assignNursingHandler struct {
	cmdRepo AppointmentCommandRepo
}

func NewAssignNursingHandler(cmdRepo AppointmentCommandRepo) *assignNursingHandler {
	return &assignNursingHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *assignNursingHandler) Handle(ctx context.Context, nursingId uuid.UUID, entity *appointmentdomain.Appointment) error {
	updateEntity, _ := appointmentdomain.NewAppointment(
		entity.GetID(),
		entity.GetServiceID(),
		entity.GetCusPackageID(),
		entity.GetPatientID(),
		&nursingId,
		entity.GetPatientAddress(),
		entity.GetPatientLatLng(),
		entity.GetStatus(),
		entity.GetTotalEstDuration(),
		entity.GetEstDate(),
		entity.GetActDate(),
		entity.GetCreatedAt(),
	)
	if err := h.cmdRepo.UpdateAppointment(ctx, updateEntity); err != nil {
		return common.NewBadRequestError().
			WithReason(fmt.Sprintf("cannot assign nursing to appointment: %v", err))
	}
	return nil
}
