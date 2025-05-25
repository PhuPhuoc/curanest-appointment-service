package appointmentqueries

import "context"

type getItemDashboardHandler struct {
	queryRepo AppointmentQueryRepo
}

func NewGetItemDashboardHandler(queryRepo AppointmentQueryRepo) *getItemDashboardHandler {
	return &getItemDashboardHandler{
		queryRepo: queryRepo,
	}
}

func (h *getItemDashboardHandler) Handle(ctx context.Context) (*AppointmentDTO, error) {
	return nil, nil
}
