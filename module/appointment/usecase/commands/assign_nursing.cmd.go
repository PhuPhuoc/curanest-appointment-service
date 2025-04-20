package apppointmentcommands

import (
	"context"

	"github.com/google/uuid"

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

func (h *assignNursingHandler) Handle(ctx context.Context, nursingId uuid.UUID, entity appointmentdomain.Appointment) error {
	return nil
}
