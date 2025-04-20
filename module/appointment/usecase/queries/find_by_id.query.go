package appointmentqueries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

type findAppointmentByIdHandler struct {
	queryRepo AppointmentQueryRepo
}

func NewFindAppointmentByIdHandler(queryRepo AppointmentQueryRepo) *findAppointmentByIdHandler {
	return &findAppointmentByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findAppointmentByIdHandler) Handle(ctx context.Context, appointmentId uuid.UUID) (*AppointmentDTO, error) {
	entity, err := h.queryRepo.FindById(ctx, appointmentId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason(fmt.Sprintf("cannot found appointment with id: %v", appointmentId.String())).
			WithInner(err.Error())
	}

	return toAppointmentDTO(entity), nil
}
