package appointmentqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

type getAppointmentsHandler struct {
	queryRepo AppointmentQueryRepo
}

func NewGetAppointmentsHandler(queryRepo AppointmentQueryRepo) *getAppointmentsHandler {
	return &getAppointmentsHandler{
		queryRepo: queryRepo,
	}
}

func (h *getAppointmentsHandler) Handle(ctx context.Context, filter *FilterGetAppointmentDTO) ([]AppointmentDTO, error) {
	if filter.Paging != nil {
		filter.Paging.Process()
	}
	entities, err := h.queryRepo.GetAppointment(ctx, filter)
	if err != nil {
		return []AppointmentDTO{}, common.NewInternalServerError().
			WithReason("cannot get appointment").
			WithInner(err.Error())
	}
	if len(entities) == 0 {
		return []AppointmentDTO{}, nil
	}

	dtos := make([]AppointmentDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = *toAppointmentDTO(&entity)
	}

	return dtos, nil
}
